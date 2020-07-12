package sql

import (
	"github.com/vietta-net/agokit/helper"
	"github.com/vietta-net/agokit/errors"
	"github.com/vietta-net/agokit/pb"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"math"
	"strings"
	"time"


	"google.golang.org/protobuf/types/known/timestamppb"

)

//Dynamic filter with multiple field with single value
func BuildWhereFilter(resultOrm *gorm.DB, filter interface{}) ( *gorm.DB, error){
	var data = make(map[string]interface{})
	//Convert Struct to map
	err := helper.Convert(filter,  &data)

	if err != nil {
		return resultOrm, errors.E(codes.Unknown, "Convert", err)
	}

	for field, value := range data {
		//Convert to lower case to map with db field name in lower case
		field = strings.ToLower(field)

		switch value := value.(type) {
		case string:
			//Check field if it has value
			if value != "" {
				field = fmt.Sprintf("`%s` = ?", field)
				resultOrm = resultOrm.Where(field, value)
			}
		case []interface{}:
			if len(value) > 0 {
				field = fmt.Sprintf("`%s` IN (?)", field)
				resultOrm = resultOrm.Where(field, value)
			}
		}
	}

	return resultOrm, nil
}

//Dynamic Search in multiple fields
func BuildWhereSearch(searchFields []string, keyword string) (string, []interface{}){
	var args   = []interface{}{}
	var whereString = ""
	//Nothing to do if searchFields and keyword are empty
	if len(searchFields) == 0 || keyword == "" {
		return whereString, args
	}

	//make where Or in (logic 1 OR logic 2 OR ...)
	var beginWhere = "("
	var endWhere 	= ")"
	var where  = make([]string, len(searchFields))

	//Search by contain a keyword
	key := "%"+ keyword + "%"
	for i, field := range searchFields {
		args = append(args, key)
		where[i] = fmt.Sprintf("%s LIKE ?", strings.ToLower(field) )
	}

	//Join a where string
	whereString = fmt.Sprintf("%s%s%s", beginWhere, strings.Join(where, " OR "),endWhere)
	return whereString, args
}

func GetPagination(resultOrm *gorm.DB, currentPage uint32, limit uint32) (interface{}, error){
	if currentPage < 1 {
		currentPage = 1
	}
	p := &Pagination{
		Count : 0,
		Pages: 0,
		Page: currentPage,
		Limit: limit,
	}
	//Count Rows Total
	err := resultOrm.Count(&p.Count).Error
	if p.Count > 0{
		pageTotal := uint32(math.Ceil(float64(p.Count / limit)))
		if pageTotal < 1{
			p.Pages = 1
		}else{
			p.Pages = pageTotal
		}
	}
	return p, err
}


func BuildWhereDateRange(resultOrm *gorm.DB, d *pb.DateRange, timezone string) ( *gorm.DB, error) {
	var err error = nil
	var dateFrom = ""
	var dateTo = ""
	if d == nil || (d.From ==nil &&  d.To == nil) {
		return resultOrm, nil
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return resultOrm, errors.E(codes.Unknown, "Timezone", err)
	}

	if d.From == nil {
		d.From = timestamppb.New(time.Unix(0,0))
	}
	if d.To == nil {
		d.To = timestamppb.New(time.Now().UTC())
	}

	if d.To.AsTime().IsZero() && d.From.AsTime().IsZero() {
		return resultOrm, nil
	}

	if d.To.AsTime().IsZero() {
		d.To = timestamppb.New(time.Now().UTC())
	}

	if d.From.AsTime().Unix() > d.To.AsTime().Unix() {
		return resultOrm, errors.E(
			codes.InvalidArgument, "Date",
			map[string]string{
				"date":"Date From should be less than or equal Date To",
			},
		)
	}

	dateFrom = d.From.AsTime().In(loc).String()
	//Add one day
	dateTo 	 = d.To.AsTime().In(loc).Add(time.Hour * 24).String()

	resultOrm = resultOrm.Where("`created_at` >= ? AND `created_at` < ?", dateFrom, dateTo)

	return resultOrm, nil
}
package app

// import (
// 	"reflect"
// )

// // MakeChan :
// // interface{}
// func MakeChan(u interface{}) []interface{} {
// 	// chanNames := make(map[int]interface{})
// 	// chanTypes := make(map[int]interface{})

// 	// t := reflect.TypeOf(u)
// 	v := reflect.ValueOf(u)
// 	chanSlice := make([]interface{}, v.NumField())

// 	for i := 0; i < v.NumField(); i++ {
// 		if v.Field(i).CanInterface() { //判断是否为可导出字段
// 			// fmt.Println(t.Field(i).Name)
// 			// fmt.Println(t.Field(i).Type)
// 			// cName := t.Field(i).Name + "Chan"
// 			// cType := t.Field(i).Type
// 			// fmt.Println(CheckType(cType))
// 			channel := make(chan interface{})
// 			chanSlice[i] = channel
// 			// channel := make(chan c)
// 			// chanNames[i] = cName
// 			// chanTypes[i] = cType
// 			// ch := make(chan reflect.TypeOf(t))
// 			// chanSlice[i] = make(chan ch)
// 			// c := make(map[string]string)
// 			// c[chanName] = t.Field(i).Type

// 			// v.Field(i).Interface()
// 			// t.Field(i).Tag
// 			// fmt.Printf("%s %s = %v -tag:%s \n",
// 		}
// 	}
// 	// fmt.Println(chanSlice)
// 	return chanSlice
// }

// // ReturnType :
// func ReturnType(u interface{}) []interface{} {
// 	t := reflect.TypeOf(u)
// 	v := reflect.ValueOf(u)
// 	r := make([]interface{}, v.NumField())
// 	for i := 0; i < v.NumField(); i++ {
// 		if v.Field(i).CanInterface() { //判断是否为可导出字段
// 			r[i] = t.Field(i).Type
// 		}
// 	}
// 	return r
// }

// // FanIn :
// func FanIn(input1, input2 <-chan interface{}) <-chan interface{} {
// 	c := make(chan interface{})
// 	go func() {
// 		for {
// 			select {
// 			case s := <-input1:
// 				c <- s
// 			case s := <-input2:
// 				c <- s
// 			}
// 		}
// 	}()
// 	return c
// }

// 這邊可以做一個made chan的
// u := GetPkgDataByOrdersID{}
// chanSlice := app.MakeChan(u)

// t := reflect.TypeOf(u)
// v := reflect.ValueOf(u)
// for i := 0; i < v.NumField(); i++ {
// 	if v.Field(i).CanInterface() { //判断是否为可导出字段
// 		fmt.Println(t.Field(i).Name)
// 		fmt.Println(t.Field(i).Type)
// 		// v.Field(i).Interface()
// 		// t.Field(i).Tag
// 		// fmt.Printf("%s %s = %v -tag:%s \n",
// 	}
// }

// pydChan := make(chan []schema.PaymentDetails)
// ivChan := make(chan []schema.Invoices)
// icChan := make(chan []schema.InvitationCode)
// mpChan := make(chan []schema.MemberPointRedeems)
// ocrChan := make(chan []schema.OrderCreateRecords)
// omrChan := make(chan []schema.OrderModRecords)

// pydChan := make(chan []schema.PaymentDetails)
// ivChan := make(chan []schema.Invoices)
// icChan := make(chan []schema.InvitationCode)
// mpChan := make(chan []schema.MemberPointRedeems)
// ocrChan := make(chan []schema.OrderCreateRecords)
// omrChan := make(chan []schema.OrderModRecords)

// go func(ordersID []int, gdb *gorm.DB) {
// 	r := pydOperator.GetByOrdersID(ordersID, gdb)
// 	pydChan <- r
// }(ordersID, gdb)

// go func(ordersID []int, gdb *gorm.DB) {
// 	r := invoiceOperator.GetByOrdersID(ordersID, gdb)
// 	ivChan <- r
// }(ordersID, gdb)

// go func(accountID []int, gdb *gorm.DB) {
// 	r := invitationCodeOperator.GetByAccountID(accountID, gdb)
// 	icChan <- r
// }(accountID, gdb)

// go func(ordersID []int, gdb *gorm.DB) {
// 	r := memberPointRedeemsOperator.GetByOrdersID(ordersID, gdb)
// 	mpChan <- r
// }(ordersID, gdb)

// go func(ordersID []int, gdb *gorm.DB) {
// 	r := orderCreateRecordOperator.GetByOrdersID(ordersID, gdb)
// 	ocrChan <- r
// }(ordersID, gdb)

// go func(ordersID []int, gdb *gorm.DB) {
// 	r := orderModRecordOperator.GetByOrdersID(ordersID, gdb)
// 	omrChan <- r
// }(ordersID, gdb)

// res := make(chan GetPkgDataByOrdersID)
// go func() {
// 	paymentDetails := <-pydChan
// 	invoices := <-ivChan
// 	invitationCode := <-icChan
// 	memberPointRedeems := <-mpChan
// 	orderModRecords := <-omrChan
// 	orderCreateRecords := <-ocrChan
// 	res <- GetPkgDataByOrdersID{paymentDetails, invoices, invitationCode, memberPointRedeems, orderCreateRecords, orderModRecords}
// }()

// result := <-res

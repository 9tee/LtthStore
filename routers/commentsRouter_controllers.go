package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["LtthStore/controllers:CategoryController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:FoodController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:FoodController"],
        beego.ControllerComments{
            Method: "GetFood",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:FoodController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:FoodController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:FoodController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:MainController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:MenuController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:MenuController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:MenuController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:MenuController"],
        beego.ControllerComments{
            Method: "GetMenu",
            Router: "/detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:StoreController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["LtthStore/controllers:StoreController"] = append(beego.GlobalControllerRouter["LtthStore/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Comment",
            Router: "/comment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

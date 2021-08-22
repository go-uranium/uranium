package uranium

//func (ushio *Ushio) HandleCategory(ctx *fiber.Ctx) error {
//	// no database writing operations,
//	// lock is unnecessary
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//
//	tname := ctx.Params("tname")
//	category := ushio.Cache.Category(tname)
//	if category == nil {
//		return fiber.NewError(http.StatusNotFound, "Category not found.")
//	}
//
//	p := ctx.Query("p", "1")
//	page, err := strconv.Atoi(p)
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	ps := ushio.Config.PageSize
//	infos, err := ushio.Data.PostsInfoByCategory(false,
//		ps, int64(page-1)*ps, category.TID)
//	if err != nil {
//		return err
//	}
//
//	return ctx.Render("category", fiber.Map{
//		"Meta": Meta{
//			Config:      *ushio.Config,
//			CurrentPage: "Home",
//		},
//		"Categories": ushio.Cache.Categories(),
//		"Nav":        nav,
//		"Data":       infos,
//		"Category":   category,
//	})
//}

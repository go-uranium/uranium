package uranium

//func (ushio *Ushio) HandleHome(ctx *fiber.Ctx) error {
//	// no database writing operations,
//	// lock is unnecessary
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//
//	p := ctx.Query("p", "1")
//	page, err := strconv.Atoi(p)
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	indexPosts,err := ushio.Cache.IndexPostInfo(int64(page))
//	if err != nil {
//		return err
//	}
//
//	return ctx.Render("home", fiber.Map{
//		"Meta": Meta{
//			Config:      *ushio.Config,
//			CurrentPage: "Home",
//		},
//		"Categories": ushio.Cache.Categories(),
//		"Nav":        nav,
//		"Data":       indexPosts,
//	})
//}

package routers

import (
	"loaner/Deliverables/controllers"
	authmiddleware "loaner/Infrastructure/authMiddleware"
	repositories "loaner/Repositories"
	usecases "loaner/Usecases"
)

func LoanRouter() {
	// userRepo := repositories.NewUserRepository(DataBase)
	loanRouter := Router.Group("/loan", authmiddleware.AuthMiddleware())
	{
		// generate repository
		loanRepo := repositories.NewLoanRepository(DataBase)
		loanUsecase := usecases.NewLoanUsecase(loanRepo)
		loanController := controllers.NewLoanController(loanUsecase)
		
		loanRouter.POST("/apply", loanController.Apply)
		loanRouter.GET("/get/:loanID", loanController.GetLoanById)
		// group by admin
		// adminLoanRouter := loanRouter.Group("", authmiddleware.IsAdminMiddleware(userRepo))
		// {
		// 	// adminLoanRouter.GET("/getall", loanController.GetAllLoans)
		// }
		
	}
}

package interfaces

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

type UserSVCInter interface {
	SignupSVC(p *pb.UserSignup) (*pb.UserResponse, error)
	VerifySVC(p *pb.UserVerify) (*pb.UserResponse, error)
	UserLogin(p *pb.UserLogin) (*pb.UserLoginResponse, error)
	ViewPackageSvc(p *pb.UserView) (*pb.UserPackage, error)
	ForgetPasswordSVC(p *pb.UserForget_Password) (*pb.UserResponse, error)
	ForgetPasswordVerifySVC(p *pb.UserForget_PasswordVerify) (*pb.UserResponse, error)
	NewPasswordSVC(p *pb.UserNew_Password) (*pb.UserResponse, error)
	UpdateProfileSVC(p *pb.UserSignup) (*pb.UserResponse, error)
	ViewDestinationSvc(p *pb.UserView) (*pb.UserDestination, error)
	ViewActivitySvc(p *pb.UserView) (*pb.UserActivity, error)
	ViewCatagoriesSvc(p *pb.UserView) (*pb.UserCategories, error)
	SearchPackageSvc(p *pb.UserSearch) (*pb.UserPackages, error)
	TraverllerDetailSVC(p *pb.UserTravellerRequest) (*pb.UserTravellerResponse, error)
	ViewFoodMenusSvc(p *pb.UserView) (*pb.UserFoodMenus, error)
	ViewPackagesSvc(p *pb.UserView) (*pb.UserPackages, error)
	OnlinePaymentSVC(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error)
	FilterPackageSvc(p *pb.UserFilter) (*pb.UserPackages, error)
	PaymentConfirmedSVC(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponse, error)
	ViewHistorySvc(p *pb.UserView) (*pb.UserHistories, error)
	ViewBookingSvc(p *pb.UserView) (*pb.UserHistory, error)
	CancelBookingSvc(p *pb.UserView) (*pb.UserResponse, error)
	RefreshTokenSVC(p *pb.TokenData) (*pb.UserLoginResponse, error)
	UserCountSVC(p *pb.UserView) (*pb.UserCount, error)
	ViewUsersSVC(p *pb.UserView) (*pb.UserUsers, error)
	ViewUserSVC(p *pb.UserView) (*pb.UserSignup, error)
}

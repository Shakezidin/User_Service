package interfaces

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/pb"
)

type UserSVCInter interface {
	SignupSVC(p *pb.UserSignup) (*pb.UserResponce, error)
	VerifySVC(p *pb.UserVerify) (*pb.UserResponce, error)
	UserLogin(p *pb.UserLogin) (*pb.UserLoginResponce, error)
	ViewPackageSvc(p *pb.UserView) (*pb.UserPackage, error) 
	ForgetPasswordSVC(p *pb.UserforgetPassword) (*pb.UserResponce, error) 
	ForgetPasswordVerifySVC(p *pb.UserforgetPasswordVerify) (*pb.UserResponce, error)
	NewPasswordSVC(p *pb.Usernewpassword) (*pb.UserResponce, error)
	UpdateProfileSVC(p *pb.UserSignup) (*pb.UserResponce, error)
	ViewDestinationSvc(p *pb.UserView) (*pb.UserDestination, error)
	ViewActivitySvc(p *pb.UserView) (*pb.UserActivity, error)
	ViewCatagoriesSvc(p *pb.UserView) (*pb.UserCategories, error)
	SearchPackageSvc(p *pb.UserSearch) (*pb.UserPackages, error) 
	TraverllerDetailSVC(p *pb.UserTravellerRequest)(*pb.UserTravellerResponse,error)
	ViewFoodMenusSvc(p *pb.UserView)(*pb.UserFoodMenus,error)
	ViewPackagesSvc(p *pb.UserView)(*pb.UserPackages,error)
	OnlinePaymentSVC(ctx context.Context,p *pb.UserBooking)(*pb.UserOnlinePaymentResponse, error)
	FilterPackageSvc(p *pb.UserFilter)(*pb.UserPackages,error)
	PaymentConfirmedSVC(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponce, error)
}

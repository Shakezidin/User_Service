package interfaces

import (
	pb "github.com/Shakezidin/pkg/user/pb"
)

type UserSVCInter interface {
	SignupSVC(p *pb.Signup) (*pb.SignupResponce, error)
	VerifySVC(p *pb.Verify) (*pb.VerifyResponce, error)
	UserLogin(p *pb.UserLogin) (*pb.UserLoginResponce, error)
	ViewPackageSvc(p *pb.ViewPackage) (*pb.ViewPacakgeResponce, error) 
	ForgetPassword(p *pb.ForgetPassword) (*pb.ForgetPasswordResponce, error)
	ForgetPasswordVerify(p *pb.ForgetPasswordVerify)(*pb.ForgetPasswordVerifyResponce,error)
	NewPassword(p *pb.Newpassword)(*pb.Newpasswordresponce,error)
}

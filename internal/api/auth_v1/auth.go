package auth_v1

import (
	auth_system "github.com/Shemistan/uzum_auth/internal/service/auth_v1"
	pb "github.com/Shemistan/uzum_auth/pkg/auth_v1"
)

type Auth struct {
	// Нужно, что бы приложение не падало в панике,
	//если какой-то АПИ еще не реализован.
	pb.UnimplementedAuthV1Server

	AuthService auth_system.IAuthSystemService
}

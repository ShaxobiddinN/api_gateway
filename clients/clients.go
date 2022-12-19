package clients

import (
	"blogpost/api_gateway/config"
	"blogpost/api_gateway/protogen/blogpost"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author  blogpost.AuthorServiceClient
	Article blogpost.ArticleServiceClient
	Auth    blogpost.AuthServiceClient

	Conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	// Conns := make([]grpc.ClientConn)
	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// Conns=append(Conns,connAuthor)
	author := blogpost.NewAuthorServiceClient(connAuthor)

	connArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// Conns=append(Conns,connArticle)
	article := blogpost.NewArticleServiceClient(connArticle)

	connAuth, err := grpc.Dial(cfg.AuthServiceGrpcHost+cfg.AuthServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// Conns=append(Conns,connAuth)
	auth := blogpost.NewAuthServiceClient(connAuth)

	Conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Author:  author,
		Article: article,
		Auth:    auth,
		Conns:   append(Conns, connAuthor, connArticle, connAuth),
	}, nil
}

// Close...
func (c *GrpcClients) Close() {
	for _, v := range c.Conns {
		v.Close()
	}
}

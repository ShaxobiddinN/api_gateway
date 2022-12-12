package clients

import (
	"blogpost/api_gateway/config"
	"blogpost/api_gateway/protogen/blogpost"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author  blogpost.AuthorServiceClient
	Article blogpost.ArticleServiceClient
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	author := blogpost.NewAuthorServiceClient(connAuthor)

	connArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	article := blogpost.NewArticleServiceClient(connArticle)

	return &GrpcClients{
		Author:  author,
		Article: article,
	}, nil
}

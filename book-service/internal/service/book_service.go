package service

import (
	"context"

	pb "github.com/Prrost/protoFinalAP2/book-service/book"
	"github.com/Prrost/protoFinalAP2/book-service/internal/model"
	"github.com/Prrost/protoFinalAP2/book-service/internal/repository"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookServiceServer struct {
	repo *repository.BookRepo
	pb.UnimplementedBookServiceServer
}

func NewBookServiceServer(repo *repository.BookRepo) *BookServiceServer {
	return &BookServiceServer{repo: repo}
}

func (s *BookServiceServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	offset := int((req.Page - 1) * req.Size)
	books, total, err := s.repo.List(ctx, offset, int(req.Size))
	if err != nil {
		return nil, err
	}
	resp := &pb.ListBooksResponse{
		TotalPages: int32((total + int64(req.Size) - 1) / int64(req.Size)),
	}
	for _, b := range books {
		resp.Books = append(resp.Books, &pb.Book{
			Id:              b.ID.String(),
			Title:           b.Title,
			Author:          b.Author,
			Isbn:            b.ISBN,
			Description:     b.Description,
			TotalCopies:     b.TotalCopies,
			AvailableCopies: b.AvailableCopies,
			CreatedAt:       timestamppb.New(b.CreatedAt),
			UpdatedAt:       timestamppb.New(b.UpdatedAt),
		})
	}
	return resp, nil
}

func (s *BookServiceServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	b, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Book{
		Id:              b.ID.String(),
		Title:           b.Title,
		Author:          b.Author,
		Isbn:            b.ISBN,
		Description:     b.Description,
		TotalCopies:     b.TotalCopies,
		AvailableCopies: b.AvailableCopies,
		CreatedAt:       timestamppb.New(b.CreatedAt),
		UpdatedAt:       timestamppb.New(b.UpdatedAt),
	}, nil
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	b := &model.Book{
		Title:           req.Title,
		Author:          req.Author,
		ISBN:            req.Isbn,
		Description:     req.Description,
		TotalCopies:     req.TotalCopies,
		AvailableCopies: req.TotalCopies,
	}
	if err := s.repo.Create(ctx, b); err != nil {
		return nil, err
	}
	return &pb.Book{
		Id:              b.ID.String(),
		Title:           b.Title,
		Author:          b.Author,
		Isbn:            b.ISBN,
		Description:     b.Description,
		TotalCopies:     b.TotalCopies,
		AvailableCopies: b.AvailableCopies,
		CreatedAt:       timestamppb.New(b.CreatedAt),
		UpdatedAt:       timestamppb.New(b.UpdatedAt),
	}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Book, error) {
	// получаем и обновляем
	bModel, err := s.repo.Get(ctx, req.Book.Id)
	if err != nil {
		return nil, err
	}
	bModel.Title = req.Book.Title
	bModel.Author = req.Book.Author
	bModel.ISBN = req.Book.Isbn
	bModel.Description = req.Book.Description
	bModel.TotalCopies = req.Book.TotalCopies
	bModel.AvailableCopies = req.Book.AvailableCopies

	if err := s.repo.Update(ctx, bModel); err != nil {
		return nil, err
	}
	return &pb.Book{
		Id:              bModel.ID.String(),
		Title:           bModel.Title,
		Author:          bModel.Author,
		Isbn:            bModel.ISBN,
		Description:     bModel.Description,
		TotalCopies:     bModel.TotalCopies,
		AvailableCopies: bModel.AvailableCopies,
		CreatedAt:       timestamppb.New(bModel.CreatedAt),
		UpdatedAt:       timestamppb.New(bModel.UpdatedAt),
	}, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *BookServiceServer) AdjustQuantity(ctx context.Context, req *pb.AdjustQuantityRequest) (*pb.AdjustQuantityResponse, error) {
	b, err := s.repo.AdjustQuantity(ctx, req.Id, req.Delta)
	if err != nil {
		return nil, err
	}
	return &pb.AdjustQuantityResponse{
		Book: &pb.Book{
			Id:              b.ID.String(),
			Title:           b.Title,
			Author:          b.Author,
			Isbn:            b.ISBN,
			Description:     b.Description,
			TotalCopies:     b.TotalCopies,
			AvailableCopies: b.AvailableCopies,
			CreatedAt:       timestamppb.New(b.CreatedAt),
			UpdatedAt:       timestamppb.New(b.UpdatedAt),
		},
	}, nil
}

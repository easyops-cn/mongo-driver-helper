package mongo_driver_helper

//go:generate mockgen -source pmongo/client.go          -destination pmongo/mock_pmongo/client.go
//go:generate mockgen -source pmongo/collection.go      -destination pmongo/mock_pmongo/collection.go
//go:generate mockgen -source pmongo/cursor.go          -destination pmongo/mock_pmongo/cursor.go
//go:generate mockgen -source pmongo/database.go        -destination pmongo/mock_pmongo/database.go
//go:generate mockgen -source pmongo/gridfs.go          -destination pmongo/mock_pmongo/gridfs.go
//go:generate mockgen -source pmongo/index_view.go      -destination pmongo/mock_pmongo/index_view.go
//go:generate mockgen -source pmongo/single_result.go   -destination pmongo/mock_pmongo/single_result.go

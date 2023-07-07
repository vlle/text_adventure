
// func main() {
//   urlExample := "postgres://postgres:postgres@localhost:5500/rec"
//   dbpool, err := pgxpool.New(context.Background(), urlExample)
//   if err != nil {
//     fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//     os.Exit(1)
//   }
//   defer dbpool.Close()
// 
//   var title string
//   var artist string
//   _, err = dbpool.Exec(context.Background(), "insert into album (title, artist, price) values ('The Dark Side of the Moon', 'Pink Floyd', 100)")
//   if err != nil {
//     fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
//     os.Exit(1)
//   }
//   err = dbpool.QueryRow(context.Background(), "select title, artist from album where id = $1", 3).Scan(&title, &artist)
//   if err != nil {
//     fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//     os.Exit(1)
//   }
// 
//   fmt.Println(title, artist)
// }

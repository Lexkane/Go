hdfs.Source(
	f,
	"hdfs://localhost::9000/home/"
	5,
	).Filter(func(line string) bool{
	}).Map(func ( line string , ch chan string){
	}).Map(func (key string) int {
	}).Reduce( func (x int , y int ) int {
	}).Map(func(x int){
	}).Run()
    }
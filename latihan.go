package main
import "fmt"

func parkir(jAwal, mAwal, jAkhir, mAkhir int){
	var jam, menit int
	jam = jAkhir-jAwal
	menit = mAkhir - mAwal
	fmt.Print(jam, " jam ", menit, " menit ")
}

func main(){
	var jAwal, mAwal, jAkhir, mAkhir int
	fmt.Scan(&jAwal, &mAwal, &jAkhir, &mAkhir)
	parkir(jAwal, mAwal, jAkhir, mAkhir)
}
package funcs

func (cmp *Company) NewWorker1(name string) *worker {
    wrk := worker{Name: name}
    cmp.Workers = append(cmp.Workers, wrk)
    return &wrk
}


func (cmp *Company) NewWorker2(name string) *worker {
    wrk := new(worker)
    wrk.Name = name
    cmp.Workers = append(cmp.Workers, *wrk)
    return wrk 
}
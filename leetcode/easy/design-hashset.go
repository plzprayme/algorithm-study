type MyHashSet struct {
    table []int 
}

func Constructor() MyHashSet {
    return MyHashSet{}
}


func (this *MyHashSet) Add(key int)  {
    for _, v := range this.table {
        if key == v {
            return
        }
    }
    
    this.table = append(this.table, key)
}


func (this *MyHashSet) Remove(key int)  {
    for i, v := range this.table {
        if key == v {
            this.table = append(this.table[:i], this.table[i+1:]...)
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    for _, v := range this.table {
        if key == v {
            return true
        }
    }
    
    return false
}
package myCalendar

//我的日程安排表 ：
//实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
//当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
//日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end) ，实数 x 的范围为 start <= x < end 。
//实现 MyCalendar 类：MyCalendar() 初始化日历对象。boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true ，
//否则，返回 false 并且不要将该日程安排添加到日历中。可以定义一个结构体来表示日程安排，包含 start 和 end 字段，然后使用一个切片来存储所有的日程安排，
//在 book 方法中，遍历切片中的日程安排，判断是否与要添加的日程安排有重叠。
/*
示例：
输入：
["MyCalendar", "book", "book", "book"]
[[], [10, 20], [15, 25], [20, 30]]
输出：
[null, true, false, true]
*/

type Event struct {
	start int
	end   int
}

type MyCalendar struct {
	events []Event
}

func Constructor() MyCalendar {
	return MyCalendar{
		events: make([]Event, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, event := range this.events {
		if start < event.end && end > event.start {
			return false
		}
	}
	this.events = append(this.events, Event{start: start, end: end})
	return true
}

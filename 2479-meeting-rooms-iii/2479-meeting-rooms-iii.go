import (
    "fmt"
    "sort"
    "container/heap"
)

type Room struct {
    number int
    end int
}

type RoomQueue struct {
    rooms []*Room
}

func (q *RoomQueue) Len() int {
    return len(q.rooms)
}

func (q *RoomQueue) Less(i, j int) bool {
    return q.rooms[i].number < q.rooms[j].number
}

func (q *RoomQueue) Swap(i, j int) {
    q.rooms[i], q.rooms[j] =  q.rooms[j], q.rooms[i]
}

func (q *RoomQueue) Push(val any) {
    q.rooms = append(q.rooms, val.(*Room))
}

func (q *RoomQueue) Pop() any {
    if q.Len() == 0 {
        return nil
    }

    n := q.Len()
    val := q.rooms[n-1]
    q.rooms[n-1] = nil
    q.rooms = q.rooms[:n-1]
    return val
}

// func (q *RoomQueue) PopFront() (*Room, bool) {
//     if q.Len() == 0 {
//         return nil, false
//     }

//     val := q.rooms[0]
//     q.rooms[0] = nil
//     q.rooms = q.rooms[1:]
//     return val, true
// }

func (q *RoomQueue) PeekFront() (*Room, bool) {
    if q.Len() == 0 {
        return nil, false
    }

    val := q.rooms[0]
    return val, true
}

func (q RoomQueue) String() string {
	var s string
	for i := range q.Len() {
		s += fmt.Sprintf(" %p:%d:%d ", q.rooms[i], q.rooms[i].number, q.rooms[i].end)
	}
	return fmt.Sprintf("[%s]", s)
}

type UsedRoomQueue struct {
    RoomQueue
}

func (q *UsedRoomQueue) Less(i, j int) bool {
    if q.rooms[i].end == q.rooms[j].end {
        return q.rooms[i].number < q.rooms[j].number
    }
    
    return q.rooms[i].end < q.rooms[j].end
}

// Push the value, then keep the slice is in ASC ordered
//func (q *UsedRoomQueue) Push(val *Room) {
//	for i := 0; i < q.Len(); i++ {
//		if q.rooms[i].end > val.end {
//            q.rooms = append(q.rooms[:i+1], q.rooms[i:]...)
//			q.rooms[i] = val
//			return
//		}
//	}
//
//    q.rooms = append(q.rooms, val)
//}

func mostBooked(n int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0] // asc order by meeting start
    })

    rooms := make([]int, n) // meeting counter on each room
    
    // unusedRooms is a priority queue,
    // the slice must in ASC order by room number
    var unusedRooms RoomQueue
    for i := 0; i < n; i++ {
        unusedRooms.Push(&Room{number: i})
    }

    // usedRooms is a priority queue,
    // the slice must in ASC oder by meeting's end
    var usedRooms UsedRoomQueue

    m := len(meetings)
    for i := 0; i < m; i++ {
        // fmt.Println("usedRooms=\t\t", usedRooms)
    
        start, end := meetings[i][0], meetings[i][1]
        // fmt.Printf("start=%d\tend=%d\n", start, end)

        // populate unused rooms
        for {
            room, ok := usedRooms.PeekFront()
            if !ok || room.end > start {
                break
            }

            unusedRoom := heap.Pop(&usedRooms).(*Room)
            unusedRoom.end = 0
            heap.Push(&unusedRooms, unusedRoom)
        }

        // fmt.Println("unusedRooms=\t", unusedRooms)

        // have an unused room
        if _, ok := unusedRooms.PeekFront(); ok {
            room := heap.Pop(&unusedRooms).(*Room)
            rooms[room.number]++

            // insert the used room to usedRooms
            room.end = end
            heap.Push(&usedRooms, room)

        // have not any unused room, pop from used rooms and get delay
        } else {
            room := heap.Pop(&usedRooms).(*Room)
            rooms[room.number]++

            room.end = room.end + (end - start)
            heap.Push(&usedRooms, room)
        }
    }

    // fmt.Println("usedRooms=\t\t", usedRooms)
    if len(rooms) > 50 {
        fmt.Println("rooms=\t\t\t", rooms[15:29])
    }

    roomMax, roomMaxCount := 0, 0
    for room, count := range rooms {
        if count > roomMaxCount {
            roomMax = room
            roomMaxCount = count
        }
    }

    return roomMax 
}
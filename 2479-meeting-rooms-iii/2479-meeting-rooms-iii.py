import heapq

class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        meetings = sorted(meetings, key=lambda i: i[0])
        rooms = [0] * n # rooms counter
        rooms_unused = [i for i in range(n)]
        rooms_used = []

        for start, end in meetings:
            # populate the unused from from used rooms
            while len(rooms_used) > 0 and rooms_used[0][0] <= start:
                _, room = heapq.heappop(rooms_used)
                heapq.heappush(rooms_unused, room)

            # push the meeting room in used rooms
            if len(rooms_unused) == 0:
                room_end, room = heapq.heappop(rooms_used)
                end = room_end + (end - start)
                heapq.heappush(rooms_unused, room)

            room = heapq.heappop(rooms_unused)
            heapq.heappush(rooms_used, (end, room))
            rooms[room] += 1
                
        room_max = 0
        room_max_count = rooms[room_max]
        for i in range(n):
            if room_max_count < rooms[i]:
                room_max = i
                room_max_count = rooms[room_max]
        
        return room_max
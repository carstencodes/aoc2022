from enum import IntEnum, auto
from functools import cached_property
from math import fabs


class State(IntEnum):
    Empty = auto()
    InRange = auto()
    Sensor = auto()
    Beacon = auto()


    def __str__(self) -> str:
        if self == State.Empty:
            return "."
        if self == State.InRange:
            return "#"
        if self == State.Sensor:
            return "S"
        if self == State.Beacon:
            return "B"
        
        raise ValueError("Invalid State")


class Measurement:
    def __init__(self, sensor: tuple[int, int], beacon: tuple[int, int]) -> None:
        self.__sensor = sensor
        self.__beacon = beacon

    @cached_property
    def states(self) -> list[list[State]]:
        result = []
        
        for y, _ in enumerate(range(0, 2*self.distance+1)):
            result.append(self.get_states_for_line(y))
            
        return result
    
    def get_states_for_line(self, y: int) -> list[State]:
        result = []
        
        min_range_x = self.distance + 1
        max_range_x = self.distance - 1
        
        direction = + 1
                
        for i in range(0, y + 1):
            min_range_x = min_range_x - direction
            max_range_x = max_range_x + direction
            if i == self.distance:
                direction = - 1
                
        for x, _ in enumerate(range(0, self.distance * 2 + 1)):
            x0 = self.min_x + x
            y0 = self.min_y + y
            
            state = State.Empty
            
            if x >= min_range_x and x <= max_range_x:
                state = State.InRange
            
            if (y0, x0) == self.__beacon:
                state = State.Beacon
                
            if (y0, x0) == self.__sensor:
                state = State.Sensor
            
            result.append(state)
        
        return result
        
    @cached_property
    def distance(self) -> int:
        return Measurement.calculate_distance(self.__sensor, self.__beacon)
    
    @property
    def sensor(self) -> tuple[int, int]:
        return self.__sensor
    
    @property
    def beacon(self) -> tuple[int, int]:
        return self.__beacon
    
    @cached_property
    def min_x(self) -> int:
        return self.__sensor[1] - self.distance
    
    @cached_property
    def min_y(self) -> int:
        return self.__sensor[0] - self.distance
    
    @cached_property
    def max_x(self) -> int:
        return self.__sensor[1] + self.distance
    
    @cached_property
    def max_y(self) -> int:
        return self.__sensor[0] + self.distance
    
    def __repr__(self) -> str:
        state_repr = ""
        #for line in self.states:
        #    for s in line:
        #        state_repr += str(s)
        #    state_repr += "\n"
        
        return f"Sensor: ({self.__sensor[1]}, {self.__sensor[0]})\nBeacon: ({self.__beacon[1]}, {self.__beacon[0]})\nDistance: {self.distance}\n{state_repr}"

    @staticmethod
    def calculate_distance(from_point: tuple[int, int], to_point: tuple[int, int]) -> int:
        delta_y = to_point[0] - from_point[0]
        delta_x = to_point[1] - from_point[1]
        
        delta_y = int(fabs(delta_y))
        delta_x = int(fabs(delta_x))
        
        return delta_x + delta_y
    

class CaveMap:
    def __init__(self, caves: list[Measurement]) -> None:
        self.__caves = caves or []
        
    @cached_property
    def min_x(self) -> int:
        return min([x.min_x for x in self.__caves])
    
    @cached_property
    def min_y(self) -> int:
        return min([x.min_y for x in self.__caves])
    
    @cached_property
    def max_x(self) -> int:
        return max([x.max_x for x in self.__caves])
    
    @cached_property
    def max_y(self) -> int:
        return max([x.max_y for x in self.__caves])
                    
    def get_line(self, line: int) -> list[State]:
        caves = [c for c in self.__caves if line >= c.min_y and line <= c.max_y]
        
        min_x = 0
        max_x = 0
        
        for cave in caves:
            min_x = min(min_x, cave.min_x)
            max_x = max(max_x, cave.max_x)
            
        dx = max_x - min_x + 1
        
        result = [State.Empty for _ in range(0, dx)]
        
        for cave in caves:
            line_translated = line - cave.min_y
            row = cave.get_states_for_line(line_translated)
            x_offset = cave.min_x
            for x, state in enumerate(row):
                x0 = x + x_offset
                if result[x0] == State.Empty:
                    result[x0] = state
            
        return result
    
    def __repr__(self) -> str:
        return f"X: {self.min_x} -> {self.max_x} x Y: {self.min_y} -> {self.max_y}"


def parse_point(text: str) -> tuple[int, int]:
    x_assign, y_assign = tuple(text.split(", "))
    x_value, y_value = (x_assign.removeprefix("x="), y_assign.removeprefix("y="))
    
    return (int(y_value), int(x_value))


def run(text: str, line_of_interest: int) -> None:
    measurements = []
    
    for line in text.splitlines():
        data_line = line.removeprefix("Sensor at ")
        data_line = data_line.replace(": closest beacon is at ", "|")
        sensor_line, beacon_line = tuple(data_line.split("|"))
        sensor, beacon = (parse_point(sensor_line),parse_point(beacon_line))
        m = Measurement(sensor, beacon)
        print(repr(m))
        
        measurements.append(m)
        
    cave_map = CaveMap(measurements)
    print(repr(cave_map))
    
    content = cave_map.get_line(line_of_interest)
    print (len(content))
    in_ranges = [1 for s in content if s == State.InRange]
    sum_of_in_ranges = sum(in_ranges)
    
    print(sum_of_in_ranges)


SAMPLE="""Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3"""

PUZZLE="""Sensor at x=2391367, y=3787759: closest beacon is at x=2345659, y=4354867
Sensor at x=1826659, y=2843839: closest beacon is at x=1654342, y=3193298
Sensor at x=980874, y=2369046: closest beacon is at x=31358, y=2000000
Sensor at x=2916267, y=2516612: closest beacon is at x=3064453, y=2107409
Sensor at x=3304786, y=844925: closest beacon is at x=3064453, y=2107409
Sensor at x=45969, y=76553: closest beacon is at x=31358, y=2000000
Sensor at x=2647492, y=1985479: closest beacon is at x=2483905, y=2123337
Sensor at x=15629, y=2015720: closest beacon is at x=31358, y=2000000
Sensor at x=3793239, y=3203486: closest beacon is at x=3528871, y=3361675
Sensor at x=3998240, y=15268: closest beacon is at x=4731853, y=1213406
Sensor at x=3475687, y=3738894: closest beacon is at x=3528871, y=3361675
Sensor at x=3993022, y=3910207: closest beacon is at x=3528871, y=3361675
Sensor at x=258318, y=2150378: closest beacon is at x=31358, y=2000000
Sensor at x=1615638, y=1108834: closest beacon is at x=2483905, y=2123337
Sensor at x=1183930, y=3997648: closest beacon is at x=1654342, y=3193298
Sensor at x=404933, y=3377916: closest beacon is at x=1654342, y=3193298
Sensor at x=3829801, y=2534117: closest beacon is at x=3528871, y=3361675
Sensor at x=2360813, y=2494240: closest beacon is at x=2483905, y=2123337
Sensor at x=2286195, y=3134541: closest beacon is at x=1654342, y=3193298
Sensor at x=15626, y=1984269: closest beacon is at x=31358, y=2000000
Sensor at x=3009341, y=3849969: closest beacon is at x=3528871, y=3361675
Sensor at x=1926292, y=193430: closest beacon is at x=1884716, y=-881769
Sensor at x=3028318, y=3091480: closest beacon is at x=3528871, y=3361675"""

#run(SAMPLE, 10)
run(PUZZLE, 20000)
using System.Text;

namespace aoc2022.day16;

internal sealed class Valve {

    public string Name { get; init; }
    public uint MaxFlowRate { get; init; }
    public bool Stuck => this.MaxFlowRate == 0;

    public IReadOnlyCollection<string> Successors { get; init; }

    public Valve(string name, uint maxFlowRate, params string[] successors) {
        this.Name = name ?? throw new ArgumentNullException(nameof(name));
        this.MaxFlowRate = maxFlowRate;
        this.Successors = successors ?? Array.Empty<string>();
    }

    public override string ToString()
    {
        string plural = this.Successors.LongCount() > 1 ? "s" : "";
        string singular = this.Successors.LongCount() <= 1 ? "s" : "";
        string valves = string.Join(", ", this.Successors);
        return $"Valve {this.Name} has flow rate={this.MaxFlowRate}; tunnel{plural} lead{singular} to valve{plural} {valves}";
    }
}

internal sealed class ValveSet {
    public IReadOnlyCollection<Valve> Valves { get; init; }

    public ValveSet(params Valve[] valves) {
        this.Valves = valves ?? Array.Empty<Valve>();
    }

    public Valve this[string name] {
        get {
            return this.Valves.SingleOrDefault(valve => StringComparer.CurrentCulture.Equals(valve.Name, name)) ?? throw new IndexOutOfRangeException(name);
        }
    }

    public override string ToString()
    {
        StringBuilder buffer = new();
        foreach (Valve v in this.Valves) {
            buffer.AppendLine(v.ToString());
        }

        return buffer.ToString();
    }
}
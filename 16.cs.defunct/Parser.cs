namespace aoc2022.day16;

internal sealed class Parser {
    private readonly string text;

    internal Parser(string text) {
        this.text = text ?? throw new ArgumentNullException(nameof(text));
    }

    internal ValveSet Parse() {
        List<Valve> valves = new();
        foreach (string line in this.text.Split(Environment.NewLine)) {
            Valve valve = ParseValveFromLine(line);
            valves.Add(valve);
        }

        return new ValveSet(valves.ToArray());
    }

    private static Valve ParseValveFromLine(string line) {
        string[] items = line.Split(";", 2);
        string valve_descriptor = items[0];
        string tunnels = items[1];

        valve_descriptor = valve_descriptor.Replace("Valve ", string.Empty);
        int next_whitespace = valve_descriptor.Trim().IndexOf(" ");
        string valve_name = valve_descriptor[..next_whitespace];
        valve_descriptor = valve_descriptor[next_whitespace..];
        string flow_rate = valve_descriptor.Split("=", 2).Last();
        uint rate = uint.Parse(flow_rate);
        string[] next_tunnels = tunnels.TrimStart().StartsWith("tunnels")
            ? tunnels[(tunnels.LastIndexOf("valves ") + "valves ".Length)..].Split(", ")
            : new string[] { tunnels[(tunnels.LastIndexOf("valve ") + "valve ".Length)..].Trim() };

        return new Valve(valve_name, rate, next_tunnels);
    }
}
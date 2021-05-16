using System.Collections.Generic;
using System.IO;

namespace Skalka.TaskRunner
{
    public class Task
    {
        public string Name;
        public string Command;
        public Dictionary<string, string> Env;
        public string Dir;

        public Task(string name, string command, Dictionary<string, string> env,
            string dir = null)
        {
            Name = name;
            Command = command;
            Env = env;
            Dir = dir ?? Directory.GetCurrentDirectory();
        }
    }
}
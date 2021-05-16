using System;
using System.Collections.Specialized;
using System.Diagnostics;

namespace Skalka.TaskRunner
{
    public class TaskRunner : ITaskRunner
    {
        public void Run(Task task)
        {
            var info = new ProcessStartInfo("cmd.exe", $"/C \"{task.Command}\"");
            info.WorkingDirectory = info.WorkingDirectory;
            info.CreateNoWindow = true;
            info.RedirectStandardOutput = true;
            info.RedirectStandardError = true;
            info.UseShellExecute = false;
            foreach (var variable in task.Env)
            {
                info.Environment.Add(variable.Key, variable.Value);
            }

            Console.WriteLine($"Executing '{task.Name}'");
            var process = Process.Start(info);
            process.ErrorDataReceived += (sender, args) => Console.WriteLine($"[{task.Name}][Err] {args.Data}");
            process.OutputDataReceived += (sender, args) => Console.WriteLine($"[{task.Name}][Out] {args.Data}");
            process.Start();
            process.BeginOutputReadLine();
            process.BeginErrorReadLine();
            process.WaitForExit();
        }
    }
}
using System;
using System.Collections;
using System.Collections.Generic;
using System.Text.RegularExpressions;

namespace Skalka.Helpers
{
    public static class EnvironmentHelper
    {
        private static readonly Regex EnvRegEx = new Regex("\\w+=.+");

        public static Dictionary<string, string> MergeEnvironments(params List<string>[] envs)
        {
            var result = new Dictionary<string, string>();
            foreach (DictionaryEntry item in Environment.GetEnvironmentVariables())
            {
                if (item.Value != null)
                {
                    result.Add(item.Key.ToString() ?? throw new InvalidOperationException(), item.Value.ToString());
                }
            }

            foreach (var env in envs)
            {
                if (env == null)
                {
                    continue;
                }

                foreach (var variable in env)
                {
                    if (!EnvRegEx.IsMatch(variable))
                    {
                        throw new Exception($"'{variable}' is not an environment variable");
                    }

                    var keyValue = variable.Trim().Split("=");
                    result.Add(keyValue[0], keyValue[1]);
                }
            }

            return result;
        }
    }
}
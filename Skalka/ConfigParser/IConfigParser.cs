namespace Skalka.ConfigParser
{
    public interface IConfigParser
    {
        T ReadConfig<T>(string path);
    }
}
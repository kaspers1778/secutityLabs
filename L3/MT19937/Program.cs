using System;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;


namespace Cripto3
{
    class Program
    {
        static async Task Main(string[] args)
        {
           //await MT19937Crack(5);
           await BetterMT19937Crack(80);

        }

        static async Task MT19937Crack(int id)
        {
            var client = new HttpClient();
            await client.GetAsync($"http://95.217.177.249/casino/playMt?id={id}&bet=1&number={1}");
            await client.GetAsync($"http://95.217.177.249/casino/createacc?id={id}");
            var mt = new MersenneTwister();
            mt.init_genrand((ulong)DateTimeOffset.UtcNow.ToUnixTimeSeconds());
            var n = (long)mt.genrand_uint32();
            Console.WriteLine($"Bet: {n}");
            HttpResponseMessage response = await client.GetAsync($"http://95.217.177.249/casino/playMt?id={id}&bet=1&number={n}");
            response.EnsureSuccessStatusCode();
            string responseBody = await response.Content.ReadAsStringAsync();
            Console.WriteLine(responseBody);
        }
        
        static async Task BetterMT19937Crack(int id)
        {
            var client = new HttpClient();
            await client.GetAsync($"http://95.217.177.249/casino/playBetterMt?id={id}&bet=1&number={1}");
            await client.GetAsync($"http://95.217.177.249/casino/createacc?id={id}");
            var outputs = new ulong[624];
            for (int i = 0; i < 624; i++)
            {
                HttpResponseMessage response = await client.GetAsync($"http://95.217.177.249/casino/playBetterMt?id={id}&bet=1&number=1");
                response.EnsureSuccessStatusCode();
                string responseBody = await response.Content.ReadAsStringAsync();
                var realNumber =ulong.Parse(responseBody.Split(":").Last().TrimEnd('}'));
                outputs[i] = Untemper(realNumber);
            }
            var mt = new MersenneTwister(outputs);
            var n = (long)mt.genrand_uint32();
            var t = 0;
            Console.WriteLine($"Bet:{n}");
            HttpResponseMessage resp = await client.GetAsync($"http://95.217.177.249/casino/playBetterMt?id={id}&bet=1&number={n}");
            resp.EnsureSuccessStatusCode();
            string respBody = await resp.Content.ReadAsStringAsync();
            Console.WriteLine(respBody);
            t++;

        }

       static ulong Untemper(ulong num)
        {
            num ^= num >> 18;
            num ^= (num << 15) & 0xefc60000;
            
            for (var i = 0; i < 7; i++) {
                num ^= (num << 7) & 0x9d2c5680;
            }
            
            for (var i = 0; i < 3; i++) {
                num ^= (num >> 11);
            }

            return num;
        }
    }
    
    
}

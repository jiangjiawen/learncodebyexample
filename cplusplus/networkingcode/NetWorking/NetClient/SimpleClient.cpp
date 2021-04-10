//https://www.youtube.com/watch?v=UbjxGvrDrbw&ab_channel=javidx9
//https://stackoverflow.com/questions/16150267/mac-os-x-get-state-of-spacebar
//https://github.com/phracker/MacOSX-SDKs/blob/master/MacOSX10.6.sdk/System/Library/Frameworks/Carbon.framework/Versions/A/Frameworks/HIToolbox.framework/Versions/A/Headers/Events.h

#include <iostream>
#include <olc_net.h>
#include <Carbon/Carbon.h>

enum class CustomMsgTypes : uint32_t
{
    ServerAccept,
    ServerDeny,
    ServerPing,
    MessageAll,
    ServerMessage,
};



class CustomClient : public olc::net::client_interface<CustomMsgTypes>
{
public:
    void PingServer()
    {
        olc::net::message<CustomMsgTypes> msg;
        msg.header.id = CustomMsgTypes::ServerPing;

        std::chrono::system_clock::time_point timeNow = std::chrono::system_clock::now();

        msg << timeNow;
        Send(msg);
    }

    void MessageAll()
    {
        olc::net::message<CustomMsgTypes> msg;
        msg.header.id = CustomMsgTypes::MessageAll;
        Send(msg);
    }
};

Boolean isPressed( unsigned short inKeyCode )
{
    unsigned char keyMap[16];
    GetKeys((BigEndianUInt32*) &keyMap);
    return (0 != ((keyMap[ inKeyCode >> 3] >> (inKeyCode & 7)) & 1));
}

int main()
{
    CustomClient c;
    c.Connect("127.0.0.1", 60000);

    bool key[3] = { false, false, false };
    bool old_key[3] = { false, false, false };

    bool bQuit = false;
    while (!bQuit)
    {
//        if (GetForegroundWindow() == GetConsoleWindow())
//        {
//            key[0] = GetAsyncKeyState('1') & 0x8000;
//            key[1] = GetAsyncKeyState('2') & 0x8000;
//            key[2] = GetAsyncKeyState('3') & 0x8000;
//        }
//ctrl+s,d,f
        key[0] = isPressed(1);
        key[1] = isPressed(2);
        key[2] = isPressed(3);

        if (key[0] && !old_key[0]) c.PingServer();
        if (key[1] && !old_key[1]) c.MessageAll();
        if (key[2] && !old_key[2]) bQuit = true;

        for (int i = 0; i < 3; i++) old_key[i] = key[i];

        if (c.IsConnected())
        {
            if (!c.Incoming().empty())
            {


                auto msg = c.Incoming().pop_front().msg;

                switch (msg.header.id)
                {
                    case CustomMsgTypes::ServerAccept:
                    {
                        // Server has responded to a ping request
                        std::cout << "Server Accepted Connection\n";
                    }
                        break;


                    case CustomMsgTypes::ServerPing:
                    {
                        // Server has responded to a ping request
                        std::chrono::system_clock::time_point timeNow = std::chrono::system_clock::now();
                        std::chrono::system_clock::time_point timeThen;
                        msg >> timeThen;
                        std::cout << "Ping: " << std::chrono::duration<double>(timeNow - timeThen).count() << "\n";
                    }
                        break;

                    case CustomMsgTypes::ServerMessage:
                    {
                        // Server has responded to a ping request
                        uint32_t clientID;
                        msg >> clientID;
                        std::cout << "Hello from [" << clientID << "]\n";
                    }
                        break;
                }
            }
        }
        else
        {
            std::cout << "Server Down\n";
            bQuit = true;
        }

    }

    return 0;
}
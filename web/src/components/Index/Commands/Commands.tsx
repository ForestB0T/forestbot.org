//@ts-ignore
import examplecommand from '../../../imgs/examplecommand.png'
//@ts-ignore
import exampledisccommand from '../../../imgs/exampledisccommand.png'

const commandList = [
    {
        cmd: '!ping',
        desc: 'See a users ping',
    },
    {
        cmd: '!bp',
        desc: 'See who has the best ping'
    },
    {
        cmd: '!fact',
        desc: 'Get a random fact'
    },
    {
        cmd: '!fact add',
        desc: 'Add a fact'
    },
    {
        cmd: '!joindate',
        desc: 'See when I first saw a user'
    },
    {
        cmd: '!lastseen',
        desc: 'Check when I lastseen a user'
    },
    {
        cmd: '!joins',
        desc: 'See how many times a user has joined',
    },
    {
        cmd: '!kd',
        desc: 'See how many kills/deaths a user has.'
    },
    {
        cmd: '!kill',
        desc: 'kill the bot'
    },
    {
        cmd: '!lastdeath',
        desc: 'see a users lastdeath'
    },
    {
        cmd: '!firstmessage',
        desc: 'see a users first message'
    },
    {
        cmd: '!lastmessage',
        desc: 'see a users last message'
    },
    {
        cmd: '!msgcount',
        desc: 'See a users message count'
    },
    {
        cmd: '!nickname',
        desc: 'Change the bots nickname (certain servers only)'
    },
    {
        cmd: '!oldnames',
        desc: 'See the old names of a user'
    },
    {
        cmd: '!playtime',
        desc: 'See how long a user has been playing'
    },
    {
        cmd: '!quote',
        desc: 'Bring up a random message'
    },
    {
        cmd: '!top kills',
        desc: 'See the top 5 users with the most kills'
    },
    {
        cmd: '!top deaths',
        desc: 'See the top 5 users with the most deaths'
    },
    {
        cmd: '!top joins',
        desc: 'See the top 5 users with the most joins/leaves'
    },
    {
        cmd: '!top playtime',
        desc: "See the top 5 users with the most playtime"
    },
    {
        cmd: '!tps',
        desc: 'See the servers current TPS (certain servers only)'
    },
    {
        cmd: '!urban',
        desc: 'Search the urban dictionary',
    },
    {
        cmd: '!wp',
        desc: 'See who has the worst ping'
    }

];

const discordCommands = [
    {
        cmd: '/tablist',
        desc: 'Get an image of the in-game Tablist'
    },
    {
        cmd: '/stats',
        desc: 'Get all statistics for the user you searched'
    },
    {
        cmd: '/livechat',
        desc: 'Setup a live chat bridge to interact with game chat via discord channel'
    },
    {
        cmd: '/setup',
        desc: 'Initialize me for the Minecraft server you use me for'
    },
    {
        cmd: '/info',
        desc: 'Get some basic info about the bot',
    }
];


const Commands = () => {
    return (
        <div className="w-full min-h-screen flex flex-col items-center justify-center py-32 gap-24" id="commands">

            <div className="bg-white flex flex-col  text-neutral-200 bg-opacity-5 w-[90%] pb-24 lg:w-2/4 min-h-[70vh] rounded-xl shadow-2xl p-8">
                <div className="flex w-full items-center justify-center h-24">
                    <h1 className="text-4xl font-bold text-center font-fredoka">Commands</h1>
                </div>

                <div className="font-poppins">
                    <ul>
                        {
                            commandList.map((command, index) => {
                                return (
                                    <li className="flex flex-col lg:flex-row lg:gap-2 gap-0">
                                        <span className="text-emerald-300 font-semibold font-mono">{command.cmd}</span>
                                        <span className="font-bold text-white">-</span>
                                        <p className="">{command.desc}</p>
                                    </li>
                                )
                            })
                        }
                    </ul>
                </div>
                <div className="mx-auto mt-9">
                    <p className="text-neutral-500 text-sm">example:</p>
                    <img src={examplecommand}></img>
                </div>

            </div>

            <div className="bg-white flex flex-col  text-neutral-200 bg-opacity-5 w-[90%] lg:w-2/4 pb-24 rounded-xl shadow-2xl p-8">
                <div className="flex w-full items-center justify-center h-24">
                    <h1 className="text-4xl font-bold text-center font-fredoka">Discord Commands</h1>
                </div>

                <div className="font-poppins">
                    <ul>
                        {
                            discordCommands.map((command, index) => {
                                return (
                                    <li className="flex flex-col lg:flex-row lg:gap-2 gap-0">
                                        <span className="text-emerald-300 font-semibold font-mono">{command.cmd}</span>
                                        <span className="font-bold text-white">-</span>
                                        <p className="">{command.desc}</p>
                                    </li>
                                )
                            })
                        }
                    </ul>
                </div>
                <div className="mx-auto mt-9">
                    <p className="text-neutral-500 text-sm">example:</p>
                    <img src={exampledisccommand}></img>
                </div>

            </div>
        </div>
    )
}

export default Commands;
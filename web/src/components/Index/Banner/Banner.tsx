//@ts-ignore
import wavey from '../../../imgs/wavey.png';

const Banner = ({ inviteUrl }: { inviteUrl: string }): JSX.Element => {
    return (
        <div className="grid grid-cols-1 lg:grid-cols-2 items-center justify-center mt-36 gap-36 py-20">

            <div className="flex flex-col w-[84%] lg:w-3/4 mx-auto items-center text-center lg:text-left" id="banner-main">
                <div className="flex flex-col gap-6 lg:w-3/4">
                    <h1 className="lg:text-7xl text-5xl font-fredoka font-extrabold text-neutral-100">Forest<span className="text-sky-500">Bot</span>.</h1>
                    <p className="lg:text-3xl text-2xl text-neutral-100 font-medium font-poppins">The <span className="text-sky-400">all in one</span> Statistics tracking utility Bot for your Minecraft server<span className="text-sky-400 font-bold">.</span></p>
                    <div className="flex flex-row gap-3 mx-auto lg:mx-0">
                        <a href={inviteUrl} target="_blank" className="border-b-8 border-b-sky-600 hover:bg-sky-700 duration-150 rounded bg-sky-500 text-white font-semibold lg:w-44 w-36 h-12 items-center flex justify-center">Invite Me</a>
                        <a href="" className="border-b-8 border-b-gray-700 hover:bg-gray-700 duration-200 rounded bg-gray-600 text-white font-semibold lg:w-44 w-36 h-12 items-center flex justify-center">Donate</a>

                    </div>
                    <img src={wavey} className="lg:-translate-x-8 px-8" />
                </div>
            </div>

            <div className="flex flex-col gap-4" id="banner-side">
                <div id="banner-info" className="flex items-center justify-center lg:justify-end lg:-translate-x-24">
                    <div className="flex justify-center gap-2 flex-col shadow-2xl bg-opacity-40 bg-gray-700 w-[80%] lg:w-2/3 min-h-[22vh] rounded-lg border-r-[12px] border-r-lime-500 p-8">
                        <h1 className="font-frodoka font-semibold text-xl text-neutral-50">What is ForestBot?</h1>
                        <p className="font-poppins text-neutral-200">ForestBot is a Minecraft bot used to track and record player statistics while having a wide variety of in-game as well as Discord commands to view and compare these stats.</p>
                    </div>
                </div>
                <div id="banner-info" className="flex items-center justify-center lg:justify-end lg:-translate-x-24">
                    <div className="flex justify-center gap-2 flex-col shadow-2xl bg-opacity-40 bg-gray-700 w-[80%] lg:w-2/3 min-h-[22vh] rounded-lg border-r-[12px] border-r-blue-500 p-8">
                        <h1 className="font-frodoka font-semibold text-xl text-neutral-50">ForestBot for your Minecraft Server!</h1>
                        <p className="font-poppins text-neutral-200">If you want Forestbot on your own Minecraft Server, a one time donation of $10 is required, although if your server is quite active with 20+ concurrent users per day, a donation is not required. </p>
                    </div>
                </div>
                <div id="banner-info" className="flex items-center justify-center lg:justify-end lg:-translate-x-24">
                    <div className="flex justify-center gap-2 flex-col shadow-2xl bg-opacity-40 bg-gray-700 w-[80%] lg:w-2/3 min-h-[22vh] rounded-lg border-r-[12px] border-r-teal-600 p-8">
                        <h1 className="font-frodoka font-semibold text-xl text-neutral-50">How does it work?</h1>
                        <p className="font-poppins text-neutral-200">Forestbot records statistics by being in-game and recording what happens. He has no ties with the server itself. He only sees what any other normal player would see, being chat messages, death messages, seeing who is online, etc.</p>
                    </div>
                </div>
            </div>
        </div>

    )
}
export default Banner;
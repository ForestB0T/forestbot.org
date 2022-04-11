
const Footer = () => {
    return (
        <div className="w-full min-h-[20vh] bg-emerald-500 flex items-center px-8 flex-col gap-28 py-24">

            <div className="flex flex-col w-full gap-20 p-14 shadow-2xl rounded bg-emerald-600 lg:w-2/4">
                <div className="font-poppins text-neutral-50 mr-auto lg:w-2/3 w-full ml-auto">
                    <h1 className="text-2xl font-semibold mb-1">Request to have ForestBot join your Minecraft server</h1>
                    <div className="flex items-center justify-start gap-3">
                        <a href="mailto:@Febzey1@gmail.com" className="px-5 py-2 bg-zinc-700 rounded font-fredoka">Email</a>
                        <p>Discord: <span className="font-semibold">Febzey#1854</span></p>
                    </div>
                </div>

                <div className="font-poppins text-neutral-50 mr-auto lg:w-2/3 w-full ml-auto">
                    <h1 className="text-2xl font-semibold mb-1">Invite ForestBot to your Discord server</h1>
                    <div className="flex items-center justify-start gap-3">
                        <a target="_blank" href="https://discord.com/api/oauth2/authorize?client_id=771280674602614825&permissions=3072&scope=bot%20applications.commands" className="px-5 py-2 bg-zinc-700 rounded font-fredoka">Invite Me</a>
                    </div>
                </div>
            </div>


        </div>
    )
}

export default Footer;
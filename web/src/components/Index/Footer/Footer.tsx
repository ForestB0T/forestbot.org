import { FaDiscord } from "react-icons/fa";

const Footer = () => {
    return (
        <div className="w-full min-h-[20vh] bg-emerald-500 flex items-center px-8 flex-col gap-28 py-24">
            <div className="">
                <ul className="flex flex-col lg:flex-row gap-2 text-center justify-center text-white font-medium text-xl">
                    <li className="hover:text-neutral-200 cursor-pointer underline underline-offset-2"><a href="https://github.com/ForestB0T">Source Code</a></li>
                    <li className="hidden lg:block"> | </li>
                    <li className="hover:text-neutral-200 cursor-pointer underline underline-offset-2"><a href="https://discord.com/invite/7Wb3PQQ2dX">Discord</a></li>
                    <li className="hidden lg:block"> | </li>
                    <li className="hover:text-neutral-200 cursor-pointer underline underline-offset-2"><a href="https://discord.com/api/oauth2/authorize?client_id=771280674602614825&permissions=3072&scope=bot%20applications.commands">Invite ForestBot</a></li>

                </ul>
            </div>


        </div>
    )
}

export default Footer;
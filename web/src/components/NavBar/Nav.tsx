import { FaDiscord } from 'react-icons/fa';

const NavBar = () => {
    return (
        <div className="absolute top-0 h-20 right-0 left-0 shadow-lg flex">
            <ul className="text-lg font-poppins text-neutral-300 flex flex-row gap-4 items-center lg:justify-between justify-center lg:mr-48 w-full">
                <div className="flex flex-row gap-4 ml-48">
                    <li className="underline-offset-[5px] underline decoration-sky-500 decoration-4 hover:underline-offset-8"><a target="_blank" href="https://discord.com/oauth2/authorize?client_id=771280674602614825&permissions=3072&scope=bot%20applications.commands">Invite Me</a></li>
                    <li className="underline-offset-[5px] underline decoration-sky-500 decoration-4 hover:underline-offset-8"><a target="_blank" href="mailto:febzey1@gmail.com">Contact</a></li>
                </div>
                <a target="_blank" href="https://discord.com/invite/7Wb3PQQ2dX"><li className="bg-[#7289da] px-5 py-1 rounded-lg "><FaDiscord className="w-9 h-9" /></li></a>
            </ul>
        </div>
    )
}

export default NavBar;
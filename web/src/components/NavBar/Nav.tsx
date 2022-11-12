import { FaDiscord } from 'react-icons/fa';

const NavBar = () => {
    return (
        <div className="absolute top-0 h-20 right-0 left-0 shadow-lg flex justify-center items-center">
            <ul className="text-lg font-poppins text-neutral-300 flex flex-row gap-4 items-center justify-center lg:justify-between lg:mr-48 w-full">
                <div className="flex flex-row gap-4 lg:ml-48 ml-8">
                    <li className="underline-offset-[5px] underline decoration-sky-500 decoration-4 hover:underline-offset-8"><a target="_blank" href="https://discord.com/oauth2/authorize?client_id=771280674602614825&permissions=3072&scope=bot%20applications.commands">Invite Me</a></li>
                    <li className="underline-offset-[5px] underline decoration-sky-500 decoration-4 hover:underline-offset-8"><a target="_blank" href="https://discord.com/invite/7Wb3PQQ2dX">Contact</a></li>
                </div>
                <a target="_blank" href="https://discord.com/invite/7Wb3PQQ2dX" className="ml-auto mr-8"><li className="bg-[#7289da] px-5 py-1 rounded-lg hover:bg-opacity-75 duration-200 "><FaDiscord className="text-4xl" /></li></a>
            </ul>
        </div>
    )
}

export default NavBar;
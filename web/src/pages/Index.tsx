import Banner from "../components/Index/Banner/Banner";
import NavBar from "../components/NavBar/Nav";

const Index = () => {

    const inviteUrl = "https://discord.com/api/oauth2/authorize?client_id=771280674602614825&permissions=3072&scope=bot%20applications.commands";

    return (
        <>
            <div className="bg-gray-800 bg-opacity-30 w-full min-h-[100vh] flex items-center flex-col lg:flex-row">
                <NavBar/>
                <Banner inviteUrl={inviteUrl} />
            </div>
        </>


    )
}

export default Index;
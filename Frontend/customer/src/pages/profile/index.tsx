import { CalendarIcon } from "lucide-react";

function Profile() {
    return (
        <div className="mx-auto space-y-4 min-w-120">
            <h1 className="font-semibold text-2xl">Профиль</h1>
            <div className="flex space-x-40">
                <div className="space-y-4">
                    <img src="https://placehold.co/200x200" alt="avatar" className="rounded-2xl"/>
                    <div className="space-y-1.5">
                        <div className="flex space-x-2 items-center">
                            <CalendarIcon size={16} className="text-gray-500"/>
                            <p>1 день</p>
                        </div>
                    </div>
                </div>
                <div className="space-y-4 w-100">
                    <div className="flex justify-between">
                        <span className="text-gray-500">Логин</span>
                        <span>buyer</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-gray-500">Фамилия</span>
                        <span>Петров</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-gray-500">Имя</span>
                        <span>Петр</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-gray-500">Номер телефона</span>
                        <span>+7 999 123 45 67</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-gray-500">Почта</span>
                        <span>buyer@example.com</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-gray-500">Телеграмм аккаунт</span>
                        <span>@buyer_petr</span>
                    </div>
                </div>
            </div>
            
        </div>
    );
}

export default Profile;
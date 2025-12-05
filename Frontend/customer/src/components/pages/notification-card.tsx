interface NotificationCardProps {
    text: string,
    date: Date,
}

export default function NotificationCard({text, date} : NotificationCardProps) {
    return (
        <div className="space-y-1 w-150">
            <h3 className="font-semibold">Уведомление</h3>
            <p className="text-wrap text-gray-500">{text}</p>
            <p>{date.toLocaleString()}</p>
        </div>
    );
}
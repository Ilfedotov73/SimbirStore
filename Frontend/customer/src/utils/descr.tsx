export function parseDesriptionToElements(desc : string) {
    const vals = desc.split(';');
    let pairs = new Map<string, string>();

    for (const val of vals) {   
        const key = val.split(':')[0]
        const value = val.split(':')[1]

        pairs = pairs.set(key, value)
    }

    return (
        <>
            {Array.from(pairs.entries()).map(([key, value]) => (
                <div className="flex justify-between" key={key}>
                    <span className="text-gray-500">{key}</span>
                    <span>{value}</span>
                </div>
            ))}
        </>
    );
}
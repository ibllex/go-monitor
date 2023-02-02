import LiquidChart from "./LiquidChart"

function DiskInfo({ data }) {
    if (!data || !data.disks) {
        return null;
    }

    return (
        <div className="px-[24px] py-[48px] mb-[24px] bg-white shadow-xs rounded-sm">
            <h2 className="text-center mb-[48px] text-2xl">{"Disk Info"}</h2>

            <div className="flex gap-x-4">
                {data.disks.map((disk, index) => {
                    let unit = 'MB';
                    let trunk = 1024 * 1024;

                    if (disk.total >= 1024 * 1024 * 1024) {
                        unit = 'GB';
                        trunk = 1024 * 1024 * 1024;
                    }

                    return (
                        <LiquidChart
                            key={index}
                            title={disk.mount_point}
                            info={`${Math.ceil(disk.used / trunk)}${unit} / ${Math.ceil(disk.total / trunk)}${unit}`}
                            percent={disk.used / disk.total}
                        />
                    );
                })}
            </div>
        </div>
    )
}

export default DiskInfo
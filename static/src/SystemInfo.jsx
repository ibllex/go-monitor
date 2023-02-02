import LiquidChart from "./LiquidChart"

function SystemInfo({data}) {

    if (!data) {
        return null;
    }

    return (
        <div className="px-[24px] py-[48px] mb-[24px] bg-white shadow-xs rounded-sm">
            <h2 className="text-center mb-[48px] text-2xl">{"System Info"}</h2>

            <div className="flex gap-x-4">
                <LiquidChart
                    title={"System Load"}
                    info={"Running Smoothly"}
                    percent={data.load/100}
                />

                <LiquidChart
                    title={"CPU Usage"}
                    info={`${data.cpu_count} Core(s)`}
                    percent={data.cpu_usage/100}
                />

                <LiquidChart
                    title={"RAM Usage"}
                    info={`${Math.ceil(data.mem_used/1024/1024)}/${Math.ceil(data.mem_total/1024/1024)}(MB)`}
                    percent={data.mem_used/data.mem_total}
                />
            </div>
        </div>
    )
}

export default SystemInfo
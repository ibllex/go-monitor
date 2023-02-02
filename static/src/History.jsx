import { Line } from '@ant-design/charts';
import { useEffect } from "react";
import { useState } from "react";
import Api from "./Api";

const data = [
    {
      "date": "01/22 14:00:05",
      "percent": 39
    },
    {
      "date": "01/22 14:10:05",
      "percent": 20
    },
    {
      "date": "01/22 14:20:05",
      "percent": 80
    },
  ];
  
const config = {
    padding: 'auto',
    xField: 'date',
    yField: 'percent',
    xAxis: {
        // type: 'timeCat',
        // tickCount: 5,
        title: {
            text: 'Date'
        }
    },
    yAxis: {
        max: 100,
        title: {
            text: 'Percent',
        }
    },
    slider: {
        start: 0,
        end: 1,
    },
};

function History() {

    const [historyInfo, setHistoryInfo] = useState(null);

    useEffect(() => {
        Api.history().then(({ data }) => {
            setHistoryInfo(data);
        });
    }, []);

    if (!historyInfo) {
        return null;
    }

    const cpu = historyInfo.map((item) => {
        return {
            "date": item.date,
            "percent": item.cpu
        };
    });

    const mem = historyInfo.map((item) => {
        return {
            "date": item.date,
            "percent": item.memory
        };
    });

    return (
        <>
            <div className="px-[24px] py-[48px] bg-white shadow-xs rounded-sm grow mb-[24px]">
                <h2 className="text-center mb-[48px] text-2xl">{"CPU"}</h2>
                <Line data={cpu} {...config} />
            </div>
            <div className="px-[24px] py-[48px] bg-white shadow-xs rounded-sm grow mb-[24px]">
                <h2 className="text-center mb-[48px] text-2xl">{"Memory"}</h2>
                <Line data={mem} {...config} />
            </div>
        </>
    );
}

export default History;

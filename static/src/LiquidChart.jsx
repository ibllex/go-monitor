import { Liquid } from '@ant-design/charts';
import { ERROR_COLOR, NORMAL_COLOR, WARNING_COLOR } from './const';

function LiquidChart({percent,title,info}) {
    const config = {
        percent,
        color: NORMAL_COLOR,
        outline: {
            border: 8,
            distance: 8,
        },
        wave: {
            length: 128,
        },
    };

    if (percent >= 0.65) {
        config.color = WARNING_COLOR;
    }

    if (percent >= 0.8) {
        config.color = ERROR_COLOR;
    }

    return (
        <div className="flex-grow flex flex-col justify-center items-center text-slate-900">
            {title && <h4 className="text-center text-lg mb-4">{title}</h4>}
            <Liquid {...config} width={200} height={200} />
            {info && <p className="text-center text-slate-600">{info}</p>}
        </div>
    );
}

export default LiquidChart;

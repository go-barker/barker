import React, { FC } from 'react';
import useSWR from 'swr';
import { fetcher } from '../fetcher';
import { VictoryPie, VictoryPieProps } from 'victory';
import { CampaignAggregatedStatistics } from 'barker-api';

function statToPieChart(
    stat: CampaignAggregatedStatistics
): { x: string; y: number; fill: string }[] {
    const {
        Delivered: delivered = 0,
        Errors: errors = 0,
        Pending: pending = 0,
        // TimedOut: timedOut = 0,
        Users: users = 0,
    } = stat;
    const left = users - errors - pending - delivered;
    return [
        {
            x: 'Left',
            y: left,
            fill: 'rgb(100%, 92.5%, 70.2%)',
        },
        {
            x: 'Delivered',
            y: delivered,
            fill: 'rgb(21.5%, 77.7%, 44.8%)',
        },
        {
            x: 'Pending',
            y: pending,
            fill: 'rgb(69%, 74.5%, 77.3%)',
        },
        {
            x: 'Error',
            y: errors,
            fill: 'rgb(82.3%, 31.8%, 31.8%)',
        },
    ].filter((d) => d.y);
}

const chartStyle: VictoryPieProps['style'] = {
    data: {
        fill: ({ datum }) => datum.fill,
    },
};

export interface CampaignStatWidgetProps {
    botID: number;
    campaignID: number;
}
export const CampaignStatWidget: FC<CampaignStatWidgetProps> = ({
    botID,
    campaignID,
}) => {
    const { data, error } = useSWR<CampaignAggregatedStatistics>(
        ['campaign.GetAggregatedStatistics', botID, campaignID],
        fetcher
    );
    return (
        <>
            {data && (
                <VictoryPie
                    data={statToPieChart(data)}
                    style={chartStyle}
                    labels={({ datum }) => `${datum.x}: ${datum.y}`}
                />
            )}
        </>
    );
};

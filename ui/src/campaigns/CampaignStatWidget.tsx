import { CampaignAggregatedStatistics } from 'barker-api';
import React, { FC } from 'react';
import useSWR from 'swr';
import { fetcher } from '../fetcher';
import { ProgressChart, ProgressChartProps } from '../ProgressChart';

function statToPieChart(
    stat: CampaignAggregatedStatistics
): ProgressChartProps['items'] {
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
            label: 'Left',
            value: left,
            color: 'rgb(100%, 92.5%, 70.2%)',
        },
        {
            label: 'Delivered',
            value: delivered,
            color: 'rgb(21.5%, 77.7%, 44.8%)',
        },
        {
            label: 'Pending',
            value: pending,
            color: 'rgb(69%, 74.5%, 77.3%)',
        },
        {
            label: 'Error',
            value: errors,
            color: 'rgb(82.3%, 31.8%, 31.8%)',
        },
    ].filter((d) => d.value);
}

export interface CampaignStatWidgetProps {
    botID: number;
    campaignID: number;
}
export const CampaignStatWidget: FC<CampaignStatWidgetProps> = ({
    botID,
    campaignID,
}) => {
    const { data } = useSWR<CampaignAggregatedStatistics>(
        ['campaign.GetAggregatedStatistics', botID, campaignID],
        fetcher
    );
    return (
        <>
            {data && (
                <ProgressChart
                    totalValue={data.Users ?? 0}
                    items={statToPieChart(data)}
                />
            )}
        </>
    );
};

import { makeStyles, Typography } from '@material-ui/core';
import React, { FC } from 'react';

interface Item {
    value: number;
    color: string;
    label: string;
}

export interface ProgressChartProps {
    items: Item[];
    totalValue: number;
}

const useStyles = makeStyles((theme) => ({
    bar: {
        height: theme.spacing(2),
        width: '100%',
        borderWidth: 1,
        borderColor: 'black',
        borderStyle: 'solid',
        display: 'flex',
    },
    legend: {
        marginTop: theme.spacing(1),
    },
    legendSquare: {
        height: theme.spacing(2),
        width: theme.spacing(2),
        borderWidth: 1,
        borderColor: 'black',
        borderStyle: 'solid',
        display: 'inline-block',
        position: 'relative',
        top: 2,
        marginRight: 2,
    },
}));

export const ProgressChart: FC<ProgressChartProps> = ({
    items,
    totalValue,
}) => {
    const classes = useStyles();
    const percentedItems: (Item & {
        percent: number;
    })[] = items.map((item) => ({
        ...item,
        percent: (item.value / totalValue) * 100,
    }));
    return (
        <div>
            <div className={classes.bar}>
                {percentedItems.map((item, i) => (
                    <div
                        key={i}
                        style={{
                            width: `${item.percent.toFixed(2)}%`,
                            backgroundColor: item.color,
                            display: 'inline-block',
                            height: '100%',
                        }}
                        title={`${item.label}: ${item.value}`}
                    ></div>
                ))}
            </div>
            <div className={classes.legend}>
                {percentedItems.map((item, i) => (
                    <Typography key={i} variant="body1">
                        <div
                            className={classes.legendSquare}
                            style={{
                                backgroundColor: item.color,
                            }}
                        ></div>
                        {item.label}: {item.value}
                    </Typography>
                ))}
            </div>
        </div>
    );
};

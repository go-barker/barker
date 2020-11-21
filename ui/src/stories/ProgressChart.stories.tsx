import React from 'react';
import { Story, Meta } from '@storybook/react/types-6-0';
import { ProgressChart, ProgressChartProps } from '../ProgressChart';

export default {
    title: 'ProgressChart',
    component: ProgressChart,
} as Meta;

const Template: Story<ProgressChartProps> = (args) => <ProgressChart {...args} />;

export const One = Template.bind({});
One.args = {
    items: [
        { value: 50, color: 'green', label: 'Green' },
        { value: 10, color: 'red', label: 'Red' }
    ],
    totalValue: 100,
};

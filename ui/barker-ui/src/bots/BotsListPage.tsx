import {
    Grid,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
} from '@material-ui/core';
import { Bot, PaginatorResponse } from 'barker-api';
import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import { Pagination } from '../Pagination';
import BotsListAppBar from './BotsListAppBar';
import { createListView } from '../createListView';

export interface BotsPageProps {
    items?: Bot[];
    error?: any;
    paging?: PaginatorResponse;
}

const BotsListView = createListView<Bot>({
    renderHeader: () => (
        <>
            <TableCell>ID</TableCell>
            <TableCell width="100%">Title</TableCell>
        </>
    ),
    renderRow: (item) => (
        <>
            <TableCell>{item.ID}</TableCell>
            <TableCell>
                <Link to={`/bots/${item.ID}`}>{item.Title}</Link>
            </TableCell>
        </>
    ),
});

export const BotsListPage: FC<BotsPageProps> = ({ items, error, paging }) => {
    return (
        <Grid container>
            <BotsListAppBar />
            <BotsListView items={items} error={error} paging={paging} />
        </Grid>
    );
};

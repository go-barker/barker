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

export interface BotsPageProps {
    items?: Bot[];
    error?: any;
    paging?: PaginatorResponse;
}

export const BotsPage: FC<BotsPageProps> = ({ items, error, paging }) => {
    if (error) return <div>failed to load</div>;
    if (!items) return <div>loading...</div>;
    return (
        <Grid container>
            <BotsListAppBar />
            <TableContainer>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell width="100%">Title</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {items.map((item, i) => (
                            <TableRow key={i}>
                                <TableCell>{item.ID}</TableCell>
                                <TableCell>
                                    <Link to={`/bots/${item.ID}`}>
                                        {item.Title}
                                    </Link>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            {paging && <Pagination paging={paging} />}
        </Grid>
    );
};

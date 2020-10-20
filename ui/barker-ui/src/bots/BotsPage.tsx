import { Bot, PaginatorResponse } from 'barker-api';
import React, { useState } from 'react';
import useSWR from 'swr';
import { fetcher } from '../fetcher';
import BotsListAppBar from './BotsListAppBar';
import {
    Grid,
    Typography,
    Table,
    TableHead,
    TableCell,
    TableRow,
    TableBody,
} from '@material-ui/core';
import { Link } from 'react-router-dom';

function BotsPage() {
    const [page, setPage] = useState(1);
    const [size, setSize] = useState(10);

    const { data: [bots, paging] = [], error } = useSWR<
        [Bot[], PaginatorResponse]
    >(['bot.List', size, page], fetcher);

    if (error) return <div>failed to load</div>;
    if (!bots) return <div>loading...</div>;
    return (
        <Grid container>
            <BotsListAppBar />

            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>ID</TableCell>
                        <TableCell>Title</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {bots.map((bot, i) => (
                        <TableRow key={i}>
                            <TableCell>{bot.ID}</TableCell>
                            <TableCell>
                                <Link to={`/bots/${bot.ID}`}>{bot.Title}</Link>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Grid>
    );
}

export default BotsPage;

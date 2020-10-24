import {
    Grid,
    Table,
    TableBody,
    TableContainer,
    TableHead,
    TableRow,
} from '@material-ui/core';
import { Bot, Campaign, PaginatorResponse, User } from 'barker-api';
import React, { FC, ReactElement } from 'react';
import { Pagination } from './Pagination';

export interface ListPageProps<Entity> {
    items?: Entity[];
    error?: any;
    paging?: PaginatorResponse;
}

export interface CreateListPageOptions<Entity> {
    renderHeader(): ReactElement;
    renderRow(item: Entity): ReactElement;
}

export function createListView<Entity extends Bot | User | Campaign>(
    options: CreateListPageOptions<Entity>
) {
    const ListPage: FC<ListPageProps<Entity>> = ({ items, error, paging }) => {
        if (error) return <div>failed to load</div>;
        if (!items) return <div>loading...</div>;
        return (
            <Grid container>
                <TableContainer>
                    <Table>
                        <TableHead>
                            <TableRow>{options.renderHeader()}</TableRow>
                        </TableHead>
                        <TableBody>
                            {items.map((item, i) => (
                                <TableRow key={i}>
                                    {options.renderRow(item)}
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                {paging && <Pagination paging={paging} />}
            </Grid>
        );
    };

    return ListPage;
}

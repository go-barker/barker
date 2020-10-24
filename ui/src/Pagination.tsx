import { TablePagination } from '@material-ui/core';
import qs from 'querystring';
import React, { FC, useCallback } from 'react';
import { useHistory } from 'react-router-dom';
import { useQuery } from './useQuery';
import { PaginatorResponse } from 'barker-api';

export interface PaginationProps {
    paging: PaginatorResponse;
}

export const Pagination: FC<PaginationProps> = ({ paging }) => {
    const history = useHistory();
    const query = useQuery();

    const handleChangePage = useCallback((event, page) => {
        history.push({
            ...query,
            search: qs.stringify({ page: page + 1 }),
        });
    }, []);

    const handleChangeRowsPerPage = useCallback((event) => {
        history.push({
            search: qs.stringify({
                ...query,
                page: 1,
                size: parseInt(event.target.value, 10),
            }),
        });
    }, []);

    return (
        <TablePagination
            component="div"
            count={paging?.TotalItems ?? 0}
            rowsPerPage={paging?.Size ?? 0}
            page={(paging?.Page ?? 1) - 1}
            onChangePage={handleChangePage}
            onChangeRowsPerPage={handleChangeRowsPerPage}
        />
    );
};

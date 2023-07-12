// Code generated by EntKit. DO NOT EDIT.
// ---------------------------------------------------------
//
// Copyright (C) 2023 EntKit. All Rights Reserved.
//
// This code is part of the EntKit library and is generated
// automatically to ensure optimal functionality and maintainability.
// Any changes made directly to this file may be overwritten
// by future code generation, leading to unexpected behavior.
//
// Please refer to the EntKit documentation for instructions on
// how to modify the library, extend its functionality or contribute
// to the project: https://entkit.com
// ---------------------------------------------------------

import React, { useState } from "react";
import { HttpError } from "@refinedev/core";
import * as RA from "@refinedev/antd";
import * as Antd from "antd";
import * as AntdIcons from "@ant-design/icons";
import * as Interfaces from "./typedefs";
import { Cursors } from "./data-provider";
import * as Custom from "./custom";
import * as View from "./view";
import * as Action from "./action";

export type GameTableProps = Antd.TableProps<Interfaces.JeppGameInterface> & {
    extendTable?: RA.useTableProps<
        Interfaces.JeppGameInterface,
        HttpError,
        any,
        Interfaces.JeppGameInterface
    >;
};
export const GameTable: React.FC<GameTableProps> = ({
    extendTable,
    ...props
}) => {
    const [cursors, setCursors] = useState<Cursors>({ first: 10 });
    const [perPage, setPerPage] = useState<number>(10);
    const table = RA.useTable<Interfaces.JeppGameInterface>({
        resource: "game",
        initialSorter: [
            {
                field: "show",
                order: "asc",
            },
        ],
        initialFilter: [
            {
                field: "id",
                value: null,
                operator: "eq",
            },
            {
                field: "show",
                value: null,
                operator: "eq",
            },
            {
                field: "airdate",
                value: null,
                operator: "eq",
            },
            {
                field: "tapedate",
                value: null,
                operator: "eq",
            },
        ],
        metaData: {
            fields: [
                "id",
                "show",
                "airdate",
                "tapedate",
                {
                    season: ["id", "number", "startdate", "enddate"],
                },
            ],
            cursors,
        },
        hasPagination: true,
        ...extendTable,
    });

    const data = table.tableQueryResult.data as any;

    return (
        <>
            <Antd.Table
                {...table.tableProps}
                pagination={false}
                rowKey="id"
                {...props}
            >
                {/* region Fields */}
                <Antd.Table.Column
                    dataIndex="id"
                    title="Id"
                    render={(value) => {
                        return <View.JeppStringViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "id",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="show"
                    title="Show"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppNumberViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "show",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="airdate"
                    title="AirDate"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppDateViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "airDate",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="tapedate"
                    title="TapeDate"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppDateViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "tapeDate",
                        table.sorter,
                    )}
                />
                {/* endregion */}

                {/* region Edges */}
                <Antd.Table.Column
                    dataIndex="season"
                    title="Season"
                    render={(value) => <View.SeasonBadge {...value} />}
                />
                {/* endregion Edges*/}

                <Antd.Table.Column<Interfaces.JeppGameInterface>
                    title="Actions"
                    dataIndex="actions"
                    render={(_, record) => (
                        <Antd.Space>
                            <Action.GameShowAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                            <Action.GameEditAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                            <Action.GameDeleteAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                        </Antd.Space>
                    )}
                />
            </Antd.Table>

            <Antd.Space style={{ marginTop: 20 }}>
                <Antd.Typography.Text type="secondary">
                    Total {data?.total || 0}
                </Antd.Typography.Text>
                <Antd.Button
                    disabled={!Boolean(data?.pageInfo?.hasPreviousPage)}
                    onClick={() => {
                        setCursors((ov) => ({
                            ...ov,
                            before: data?.pageInfo?.startCursor,
                            last: perPage,
                            after: undefined,
                            first: undefined,
                        }));
                    }}
                >
                    <AntdIcons.LeftOutlined />
                    Prev
                </Antd.Button>
                <Antd.Button
                    disabled={!Boolean(data?.pageInfo?.hasNextPage)}
                    onClick={() => {
                        setCursors((ov) => {
                            return {
                                ...ov,
                                after: data?.pageInfo?.endCursor,
                                first: perPage,
                                before: undefined,
                                last: undefined,
                            };
                        });
                    }}
                >
                    Next
                    <AntdIcons.RightOutlined />
                </Antd.Button>
                <Antd.Select
                    labelInValue
                    defaultValue={{ value: 10, label: "10 / page" }}
                    style={{ width: 110 }}
                    onChange={(value) => {
                        setPerPage(value.value);
                        setCursors((ov) => ({
                            ...ov,
                            // Return to first page
                            first: value.value,
                            last: undefined,
                            before: undefined,
                            after: undefined,
                        }));
                    }}
                    options={[
                        { value: 10, label: "10 / page" },
                        { value: 20, label: "20 / page" },
                        { value: 50, label: "50 / page" },
                        { value: 100, label: "100 / page" },
                    ]}
                />
            </Antd.Space>
        </>
    );
};
export type SeasonTableProps =
    Antd.TableProps<Interfaces.JeppSeasonInterface> & {
        extendTable?: RA.useTableProps<
            Interfaces.JeppSeasonInterface,
            HttpError,
            any,
            Interfaces.JeppSeasonInterface
        >;
    };
export const SeasonTable: React.FC<SeasonTableProps> = ({
    extendTable,
    ...props
}) => {
    const [cursors, setCursors] = useState<Cursors>({ first: 10 });
    const [perPage, setPerPage] = useState<number>(10);
    const table = RA.useTable<Interfaces.JeppSeasonInterface>({
        resource: "season",
        initialSorter: [
            {
                field: "number",
                order: "asc",
            },
        ],
        initialFilter: [
            {
                field: "id",
                value: null,
                operator: "eq",
            },
            {
                field: "number",
                value: null,
                operator: "eq",
            },
            {
                field: "startdate",
                value: null,
                operator: "eq",
            },
            {
                field: "enddate",
                value: null,
                operator: "eq",
            },
        ],
        metaData: {
            fields: [
                "id",
                "number",
                "startdate",
                "enddate",
                {
                    games: [
                        /*{
                        edges: [
                            {
                                node: [
                                    "id",
                                    "show",
                                    "airdate",
                                    "tapedate",
                                ]
                            },
                        ],
                    },*/
                        "totalCount",
                    ],
                },
            ],
            cursors,
        },
        hasPagination: true,
        ...extendTable,
    });

    const data = table.tableQueryResult.data as any;

    return (
        <>
            <Antd.Table
                {...table.tableProps}
                pagination={false}
                rowKey="id"
                {...props}
            >
                {/* region Fields */}
                <Antd.Table.Column
                    dataIndex="id"
                    title="Id"
                    render={(value) => {
                        return <View.JeppStringViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "id",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="number"
                    title="Number"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppNumberViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "number",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="startdate"
                    title="StartDate"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppDateViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "startDate",
                        table.sorter,
                    )}
                />
                <Antd.Table.Column
                    dataIndex="enddate"
                    title="EndDate"
                    sorter={{}}
                    render={(value) => {
                        return <View.JeppDateViewOnList value={value} />;
                    }}
                    filterDropdown={(props) => (
                        <RA.FilterDropdown {...props}>
                            <Antd.Input />
                        </RA.FilterDropdown>
                    )}
                    defaultSortOrder={RA.getDefaultSortOrder(
                        "endDate",
                        table.sorter,
                    )}
                />
                {/* endregion */}

                {/* region Edges */}
                <Antd.Table.Column
                    dataIndex="games"
                    title="Games"
                    render={(value) => (
                        <span>{value?.totalCount || "No"} Items</span>
                    )}
                />
                {/* endregion Edges*/}

                <Antd.Table.Column<Interfaces.JeppSeasonInterface>
                    title="Actions"
                    dataIndex="actions"
                    render={(_, record) => (
                        <Antd.Space>
                            <Action.SeasonShowAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                            <Action.SeasonEditAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                            <Action.SeasonDeleteAction
                                recordItemIDs={[record.id]}
                                size="small"
                                hideText={true}
                            />
                        </Antd.Space>
                    )}
                />
            </Antd.Table>

            <Antd.Space style={{ marginTop: 20 }}>
                <Antd.Typography.Text type="secondary">
                    Total {data?.total || 0}
                </Antd.Typography.Text>
                <Antd.Button
                    disabled={!Boolean(data?.pageInfo?.hasPreviousPage)}
                    onClick={() => {
                        setCursors((ov) => ({
                            ...ov,
                            before: data?.pageInfo?.startCursor,
                            last: perPage,
                            after: undefined,
                            first: undefined,
                        }));
                    }}
                >
                    <AntdIcons.LeftOutlined />
                    Prev
                </Antd.Button>
                <Antd.Button
                    disabled={!Boolean(data?.pageInfo?.hasNextPage)}
                    onClick={() => {
                        setCursors((ov) => {
                            return {
                                ...ov,
                                after: data?.pageInfo?.endCursor,
                                first: perPage,
                                before: undefined,
                                last: undefined,
                            };
                        });
                    }}
                >
                    Next
                    <AntdIcons.RightOutlined />
                </Antd.Button>
                <Antd.Select
                    labelInValue
                    defaultValue={{ value: 10, label: "10 / page" }}
                    style={{ width: 110 }}
                    onChange={(value) => {
                        setPerPage(value.value);
                        setCursors((ov) => ({
                            ...ov,
                            // Return to first page
                            first: value.value,
                            last: undefined,
                            before: undefined,
                            after: undefined,
                        }));
                    }}
                    options={[
                        { value: 10, label: "10 / page" },
                        { value: 20, label: "20 / page" },
                        { value: 50, label: "50 / page" },
                        { value: 100, label: "100 / page" },
                    ]}
                />
            </Antd.Space>
        </>
    );
};

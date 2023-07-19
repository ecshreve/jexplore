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

import { useState, useEffect } from "react";
import { Typography, Avatar, AutoComplete, Input } from "antd";
import { useList, useLink } from "@refinedev/core";
import * as Interfaces from "./typedefs";
import debounce from "lodash/debounce";
import { FileImageOutlined, SearchOutlined } from "@ant-design/icons";

interface IOptionGroup {
    value: string;
    label: string | React.ReactNode;
}

interface IOptions {
    label: string | React.ReactNode;
    options: IOptionGroup[];
}

const { Text } = Typography;

export const SearchComponent: React.FC = () => {
    const [options, setOptions] = useState<IOptions[]>([]);
    const [value, setValue] = useState<string>("");
    const Link = useLink();

    const renderTitle = (title: string) => (
        <Typography.Title>
            <Typography.Text style={{ fontSize: "16px" }}>
                {title}
            </Typography.Text>
        </Typography.Title>
    );

    const renderItem = (
        title: string,
        imageUrl: string | null,
        link: string,
    ) => ({
        value: title,
        key: link,
        label: (
            <Link
                key={title + link}
                to={link}
                style={{ display: "flex", alignItems: "center" }}
            >
                {imageUrl ? (
                    <Avatar
                        size={48}
                        src={imageUrl}
                        style={{ minWidth: "48px" }}
                    />
                ) : (
                    <FileImageOutlined style={{ fontSize: "48px" }} />
                )}
                <Text style={{ marginLeft: "16px" }}>{title}</Text>
            </Link>
        ),
    });
    const { refetch: refetchCategory } =
        useList<Interfaces.JeppCategoryInterface>({
            resource: "category",
            metaData: {
                fields: ["id", "name"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.name),
                            null,
                            window.environment.appPath +
                                "category/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Category"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchClue } = useList<Interfaces.JeppClueInterface>({
        resource: "clue",
        metaData: {
            fields: ["id", "id"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.id),
                        null,
                        window.environment.appPath +
                            "clue/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Clue"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchGame } = useList<Interfaces.JeppGameInterface>({
        resource: "game",
        metaData: {
            fields: ["id", "id"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.id),
                        null,
                        window.environment.appPath +
                            "game/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Game"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchSeason } = useList<Interfaces.JeppSeasonInterface>({
        resource: "season",
        metaData: {
            fields: ["id", "number"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.number),
                        null,
                        window.environment.appPath +
                            "season/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Season"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });

    useEffect(() => {
        setOptions([]);
        if (value.length < 3) {
            return;
        }
        refetchCategory();
        refetchClue();
        refetchGame();
        refetchSeason();
    }, [value]);

    return (
        <AutoComplete
            style={{
                width: "100%",
                maxWidth: "550px",
            }}
            options={options}
            filterOption={false}
            onSearch={debounce((value: string) => setValue(value), 300)}
        >
            <Input
                size="large"
                placeholder="Search"
                suffix={<SearchOutlined />}
            />
        </AutoComplete>
    );
};
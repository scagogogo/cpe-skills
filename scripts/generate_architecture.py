#!/usr/bin/env python3
"""
Generate an architecture diagram for cpe-skills project.
Shows the layered architecture and data flow between modules.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch, FancyArrowPatch
import matplotlib.font_manager as fm
import numpy as np

# Register CJK font
_CJK_FONT = None
for fname in ['Noto Sans CJK SC', 'AR PL UMing CN', 'WenQuanYi Micro Hei']:
    matches = [f for f in fm.findSystemFonts() if fname.lower().replace(' ', '') in f.lower().replace(' ', '')]
    if matches:
        _CJK_FONT = fname
        break

if _CJK_FONT is None:
    for f in fm.fontManager.ttflist:
        if 'CJK' in f.name or 'cjk' in f.name.lower():
            _CJK_FONT = f.name
            break

FONT_FAMILY = [_CJK_FONT, 'DejaVu Sans'] if _CJK_FONT else ['DejaVu Sans']


def draw_architecture(output_path, lang="en"):
    """Draw the layered architecture diagram."""

    fig_width = 26
    fig_height = 16

    fig, ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), dpi=150)
    ax.set_xlim(0, fig_width)
    ax.set_ylim(0, fig_height)
    ax.axis('off')

    # Background
    gradient = np.linspace(0, 1, 256).reshape(1, -1)
    gradient = np.vstack([gradient] * 256)
    ax.imshow(gradient, aspect='auto', extent=[0, fig_width, 0, fig_height],
              cmap=matplotlib.colors.LinearSegmentedColormap.from_list(
                  'bg', ['#0D1117', '#161B22']), alpha=0.95, zorder=0)

    if lang == "en":
        layers = [
            {"name": "Integration Layer", "color": "#58A6FF", "y": 12.5,
             "desc": "SKILLS · Go SDK · CLI · MCP Server",
             "items": ["SKILLS", "Go SDK", "CLI", "MCP Server"]},
            {"name": "Application Layer", "color": "#3FB950", "y": 9.8,
             "desc": "SBOM Engine · Vulnerability Report · Risk Scoring · VEX · Export",
             "items": ["SBOM Engine", "Vuln Report", "Risk Scoring", "VEX Document", "Export"]},
            {"name": "Analysis Layer", "color": "#D2A8FF", "y": 7.1,
             "desc": "Matching · Reachability · Dependency Graph · License · Remediation",
             "items": ["Match Engine", "Reachability", "Dep Graph", "License", "Remediation"]},
            {"name": "Data Source Layer", "color": "#F0883E", "y": 4.4,
             "desc": "NVD · OSV · EPSS · CISA KEV · Multi-Source Search",
             "items": ["NVD Feed", "OSV API", "EPSS API", "CISA KEV", "Multi-Source"]},
            {"name": "Core Layer", "color": "#FF7B72", "y": 1.7,
             "desc": "CPE Parse/Format · WFN Binding · Validation · Storage · PURL · Ecosystem",
             "items": ["CPE Parse", "WFN Bind", "Validation", "Storage", "PURL", "Ecosystem"]},
        ]
        title = "cpe-skills Architecture"
        subtitle = "Comprehensive CPE Toolkit — Full Lifecycle from Parsing to Vulnerability Management"
    else:
        layers = [
            {"name": "集成层", "color": "#58A6FF", "y": 12.5,
             "desc": "SKILLS · Go SDK · CLI · MCP 服务器",
             "items": ["SKILLS", "Go SDK", "CLI", "MCP 服务器"]},
            {"name": "应用层", "color": "#3FB950", "y": 9.8,
             "desc": "SBOM 引擎 · 漏洞报告 · 风险评分 · VEX · 导出",
             "items": ["SBOM 引擎", "漏洞报告", "风险评分", "VEX 文档", "导出"]},
            {"name": "分析层", "color": "#D2A8FF", "y": 7.1,
             "desc": "匹配 · 可达性 · 依赖图 · 许可证 · 修复",
             "items": ["匹配引擎", "可达性", "依赖图", "许可证", "修复建议"]},
            {"name": "数据源层", "color": "#F0883E", "y": 4.4,
             "desc": "NVD · OSV · EPSS · CISA KEV · 多源搜索",
             "items": ["NVD 数据源", "OSV API", "EPSS API", "CISA KEV", "多源搜索"]},
            {"name": "核心层", "color": "#FF7B72", "y": 1.7,
             "desc": "CPE 解析/格式化 · WFN 绑定 · 校验 · 存储 · PURL · 生态系统",
             "items": ["CPE 解析", "WFN 绑定", "校验", "存储", "PURL", "生态系统"]},
        ]
        title = "cpe-skills 架构图"
        subtitle = "全面的 CPE 工具包 — 从解析到漏洞管理的全生命周期"

    layer_height = 2.0
    margin_x = 1.5
    layer_width = fig_width - 2 * margin_x

    for layer in layers:
        y = layer["y"]
        color = layer["color"]

        # Layer background with rounded corners
        layer_bg = FancyBboxPatch(
            (margin_x, y), layer_width, layer_height,
            boxstyle="round,pad=0.2", facecolor=color, edgecolor=color,
            linewidth=2, alpha=0.08, zorder=2
        )
        ax.add_patch(layer_bg)

        # Layer border
        layer_border = FancyBboxPatch(
            (margin_x, y), layer_width, layer_height,
            boxstyle="round,pad=0.2", facecolor='none', edgecolor=color,
            linewidth=1.8, alpha=0.4, zorder=3
        )
        ax.add_patch(layer_border)

        # Layer name label (left)
        name_box_w = 3.5
        name_box = FancyBboxPatch(
            (margin_x + 0.2, y + layer_height/2 - 0.35), name_box_w, 0.7,
            boxstyle="round,pad=0.1", facecolor=color, edgecolor='none',
            alpha=0.3, zorder=4
        )
        ax.add_patch(name_box)
        ax.text(margin_x + 0.2 + name_box_w/2, y + layer_height/2, layer["name"],
                ha='center', va='center', fontsize=12, fontweight='bold',
                color=color, alpha=0.95, zorder=5, fontfamily=FONT_FAMILY)

        # Item boxes
        items = layer["items"]
        n_items = len(items)
        item_start_x = margin_x + 4.5
        item_end_x = margin_x + layer_width - 0.5
        item_total_width = item_end_x - item_start_x
        item_width = item_total_width / n_items

        for i, item in enumerate(items):
            ix = item_start_x + i * item_width + item_width / 2
            iy = y + layer_height / 2

            ib_w = min(item_width - 0.4, 3.2)
            ib_h = 0.65

            item_box = FancyBboxPatch(
                (ix - ib_w/2, iy - ib_h/2), ib_w, ib_h,
                boxstyle="round,pad=0.08", facecolor=color, edgecolor='white',
                linewidth=0.8, alpha=0.2, zorder=4
            )
            ax.add_patch(item_box)

            ax.text(ix, iy, item, ha='center', va='center',
                    fontsize=9.5, color='white', alpha=0.92, zorder=5,
                    fontfamily=FONT_FAMILY)

    # Draw flow arrows between layers
    arrow_color = '#484F58'
    for i in range(len(layers) - 1):
        upper_y = layers[i]["y"]
        lower_y = layers[i+1]["y"] + layer_height
        mid_x = fig_width / 2

        # Down arrow (request flow)
        ax.annotate('', xy=(mid_x - 1.5, lower_y + 0.15), xytext=(mid_x - 1.5, upper_y - 0.15),
                     arrowprops=dict(arrowstyle='-|>', color=arrow_color, lw=1.8, alpha=0.6),
                     zorder=6)
        # Up arrow (result flow)
        ax.annotate('', xy=(mid_x + 1.5, upper_y - 0.15), xytext=(mid_x + 1.5, lower_y + 0.15),
                     arrowprops=dict(arrowstyle='-|>', color=arrow_color, lw=1.8, alpha=0.6),
                     zorder=6)

    # Flow labels
    if lang == "en":
        ax.text(fig_width / 2 - 1.5, (layers[0]["y"] + layers[-1]["y"] + layer_height) / 2,
                "Request ▼", ha='center', va='center', fontsize=8, color='#8B949E',
                rotation=90, zorder=7, fontfamily=FONT_FAMILY)
        ax.text(fig_width / 2 + 1.5, (layers[0]["y"] + layers[-1]["y"] + layer_height) / 2,
                "▲ Result", ha='center', va='center', fontsize=8, color='#8B949E',
                rotation=90, zorder=7, fontfamily=FONT_FAMILY)
    else:
        ax.text(fig_width / 2 - 1.5, (layers[0]["y"] + layers[-1]["y"] + layer_height) / 2,
                "请求 ▼", ha='center', va='center', fontsize=8, color='#8B949E',
                rotation=90, zorder=7, fontfamily=FONT_FAMILY)
        ax.text(fig_width / 2 + 1.5, (layers[0]["y"] + layers[-1]["y"] + layer_height) / 2,
                "▲ 结果", ha='center', va='center', fontsize=8, color='#8B949E',
                rotation=90, zorder=7, fontfamily=FONT_FAMILY)

    # Title
    ax.text(fig_width / 2, fig_height - 0.5, title,
            ha='center', va='top', fontsize=18, fontweight='bold',
            color='#C9D1D9', zorder=11, fontfamily=FONT_FAMILY)
    ax.text(fig_width / 2, fig_height - 1.2, subtitle,
            ha='center', va='top', fontsize=10, color='#8B949E', style='italic', zorder=11,
            fontfamily=FONT_FAMILY)

    plt.tight_layout(pad=0.5)
    plt.savefig(output_path, dpi=150, bbox_inches='tight',
                facecolor='#0D1117', edgecolor='none')
    plt.close()
    print(f"Saved architecture diagram to {output_path}")


if __name__ == "__main__":
    draw_architecture("docs/images/architecture_en.png", lang="en")
    draw_architecture("docs/images/architecture_zh.png", lang="zh")

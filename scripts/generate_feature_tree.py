#!/usr/bin/env python3
"""
Generate a feature tree mindmap for cpe-skills project.
Renders a visually appealing tree diagram showing all feature categories and sub-features.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
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

# ============================================================
# Feature tree data
# ============================================================
FEATURE_TREE = [
    ("CPE Core", "#E74C3C", [
        "CPE 2.2 Parsing",
        "CPE 2.3 Parsing",
        "Auto Format Detection",
        "CPE Struct & Fields",
        "WFN Binding (FS/URI)",
        "Character Escaping",
    ]),
    ("Matching", "#3498DB", [
        "NISTIR 7696 Matching",
        "Exact / Subset / Superset",
        "Disjoint Detection",
        "Advanced Fuzzy Match",
        "Partial / Regex Match",
        "Distance-Based Match",
        "Batch Matching",
    ]),
    ("Generation", "#2ECC71", [
        "CPE Creation",
        "Builder Pattern",
        "Fuzzy Generation",
        "Template Generation",
        "Merge & Fill Defaults",
        "Random CPE",
    ]),
    ("Validation", "#F39C12", [
        "CPE Validation",
        "Component Validation",
        "Vendor Normalization",
        "Product Normalization",
        "Version Comparison",
        "Version Range Matching",
    ]),
    ("Storage", "#9B59B6", [
        "Memory Storage",
        "File Storage + Cache",
        "Storage Manager",
        "CPE Index",
        "Search & Filter",
        "CPE Dictionary (XML)",
    ]),
    ("Vulnerability", "#E67E22", [
        "CVE Reference & Query",
        "NVD Feed Integration",
        "OSV API Client",
        "EPSS Scoring",
        "CISA KEV Catalog",
        "Vulnerability Report",
        "Risk Scoring & Priority",
    ]),
    ("SBOM", "#1ABC9C", [
        "SBOM Model & Builder",
        "CycloneDX Parser",
        "SPDX Parser",
        "SBOM Merge & Diff",
        "Component Enrichment",
        "Pedigree & Evidence",
        "SBOM Validation",
    ]),
    ("PURL & Ecosystem", "#E84393", [
        "Package URL (PURL)",
        "CPE <-> PURL Mapping",
        "Batch CPE/PURL Conv.",
        "Ecosystem Detection",
        "Ecosystem Hints",
    ]),
    ("SCA & Reach", "#00B894", [
        "Dependency Graph",
        "Reachability Analysis",
        "Batch Reachability",
        "Remediation Advice",
        "License Detection",
        "License Compliance",
    ]),
    ("Export & VEX", "#6C5CE7", [
        "JSON / CSV / SARIF",
        "CycloneDX SBOM Export",
        "SPDX SBOM Export",
        "VEX Document",
        "VEX from Findings",
    ]),
    ("Infrastructure", "#636E72", [
        "Data Source Registry",
        "Multi-Source Search",
        "Applicability Lang.",
        "Set Operations",
        "Structured Errors",
        "Logging Framework",
        "Manifest Bridge",
    ]),
]

CAT_ZH = {
    "CPE Core": "CPE 核心",
    "Matching": "匹配引擎",
    "Generation": "生成构建",
    "Validation": "校验规范化",
    "Storage": "存储索引",
    "Vulnerability": "漏洞管理",
    "SBOM": "软件物料清单",
    "PURL & Ecosystem": "PURL 与生态",
    "SCA & Reach": "SCA 与可达性",
    "Export & VEX": "导出与 VEX",
    "Infrastructure": "基础设施",
}

CHILD_ZH = {
    "CPE 2.2 Parsing": "CPE 2.2 解析",
    "CPE 2.3 Parsing": "CPE 2.3 解析",
    "Auto Format Detection": "自动格式检测",
    "CPE Struct & Fields": "CPE 结构与字段",
    "WFN Binding (FS/URI)": "WFN 绑定 (FS/URI)",
    "Character Escaping": "字符转义",
    "NISTIR 7696 Matching": "NISTIR 7696 匹配",
    "Exact / Subset / Superset": "精确/子集/超集",
    "Disjoint Detection": "不相交检测",
    "Advanced Fuzzy Match": "高级模糊匹配",
    "Partial / Regex Match": "部分/正则匹配",
    "Distance-Based Match": "距离匹配",
    "Batch Matching": "批量匹配",
    "CPE Creation": "CPE 创建",
    "Builder Pattern": "Builder 构建器",
    "Fuzzy Generation": "模糊生成",
    "Template Generation": "模板生成",
    "Merge & Fill Defaults": "合并与默认填充",
    "Random CPE": "随机 CPE",
    "CPE Validation": "CPE 校验",
    "Component Validation": "组件校验",
    "Vendor Normalization": "供应商规范化",
    "Product Normalization": "产品规范化",
    "Version Comparison": "版本比较",
    "Version Range Matching": "版本范围匹配",
    "Memory Storage": "内存存储",
    "File Storage + Cache": "文件存储+缓存",
    "Storage Manager": "存储管理器",
    "CPE Index": "CPE 索引",
    "Search & Filter": "搜索与过滤",
    "CPE Dictionary (XML)": "CPE 字典 (XML)",
    "CVE Reference & Query": "CVE 引用与查询",
    "NVD Feed Integration": "NVD 数据源集成",
    "OSV API Client": "OSV API 客户端",
    "EPSS Scoring": "EPSS 评分",
    "CISA KEV Catalog": "CISA KEV 目录",
    "Vulnerability Report": "漏洞报告",
    "Risk Scoring & Priority": "风险评分与优先级",
    "SBOM Model & Builder": "SBOM 模型与构建",
    "CycloneDX Parser": "CycloneDX 解析",
    "SPDX Parser": "SPDX 解析",
    "SBOM Merge & Diff": "SBOM 合并与对比",
    "Component Enrichment": "组件增强",
    "Pedigree & Evidence": "谱系与证据",
    "SBOM Validation": "SBOM 校验",
    "Package URL (PURL)": "包统一资源定位",
    "CPE <-> PURL Mapping": "CPE <-> PURL 映射",
    "Batch CPE/PURL Conv.": "批量 CPE/PURL 转换",
    "Ecosystem Detection": "生态系统检测",
    "Ecosystem Hints": "生态提示",
    "Dependency Graph": "依赖图",
    "Reachability Analysis": "可达性分析",
    "Batch Reachability": "批量可达性",
    "Remediation Advice": "修复建议",
    "License Detection": "许可证检测",
    "License Compliance": "许可证合规",
    "JSON / CSV / SARIF": "JSON / CSV / SARIF",
    "CycloneDX SBOM Export": "CycloneDX SBOM 导出",
    "SPDX SBOM Export": "SPDX SBOM 导出",
    "VEX Document": "VEX 文档",
    "VEX from Findings": "从发现生成 VEX",
    "Data Source Registry": "数据源注册",
    "Multi-Source Search": "多源搜索",
    "Applicability Lang.": "适用性语言",
    "Set Operations": "集合运算",
    "Structured Errors": "结构化错误",
    "Logging Framework": "日志框架",
    "Manifest Bridge": "清单桥接",
}


def draw_feature_tree(output_path, lang="en"):
    """Draw the feature tree as a mindmap with category boxes and children below."""

    root_label = "cpe-skills\n功能树" if lang == "zh" else "cpe-skills\nFeature Tree"

    n_cats = len(FEATURE_TREE)
    fig_width = 32
    fig_height = 22

    fig, ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), dpi=150)
    ax.set_xlim(-0.5, fig_width + 0.5)
    ax.set_ylim(-1, fig_height + 0.5)
    ax.axis('off')

    # Dark background
    ax.set_facecolor('#0D1117')
    fig.patch.set_facecolor('#0D1117')

    # Root node
    root_x = fig_width / 2
    root_y = fig_height - 1.5
    root_box = FancyBboxPatch(
        (root_x - 3, root_y - 0.55), 6, 1.1,
        boxstyle="round,pad=0.18", facecolor='#58A6FF', edgecolor='#79C0FF',
        linewidth=3, zorder=10
    )
    ax.add_patch(root_box)
    ax.text(root_x, root_y, root_label, ha='center', va='center',
            fontsize=22, fontweight='bold', color='white', zorder=11,
            fontfamily=FONT_FAMILY)

    # Layout: 2 rows of categories
    cols_top = 6
    cols_bot = 5
    row_y_top = fig_height - 5.0
    row_y_bot = fig_height - 10.0

    cat_positions = []
    spacing_top = fig_width / (cols_top + 1)
    for i in range(cols_top):
        cat_positions.append((spacing_top * (i + 1), row_y_top))
    spacing_bot = fig_width / (cols_bot + 1)
    for i in range(cols_bot):
        cat_positions.append((spacing_bot * (i + 1), row_y_bot))

    # Draw connections from root to categories
    for idx, (cat_name, color, children) in enumerate(FEATURE_TREE):
        cx, cy = cat_positions[idx]
        # Curved line from root bottom to category top
        ax.annotate('',
            xy=(cx, cy + 0.5),
            xytext=(root_x, root_y - 0.55),
            arrowprops=dict(
                arrowstyle='-',
                color=color,
                lw=2.2,
                connectionstyle='arc3,rad=0.12',
                alpha=0.65
            ),
            zorder=5
        )

    # Draw categories and children
    for idx, (cat_name, color, children) in enumerate(FEATURE_TREE):
        cx, cy = cat_positions[idx]

        # Category label
        cat_label = CAT_ZH.get(cat_name, cat_name) if lang == "zh" else cat_name

        # Category box
        box_w = 3.8
        box_h = 0.9
        # Shadow
        shadow = FancyBboxPatch(
            (cx - box_w/2 + 0.06, cy - box_h/2 - 0.06), box_w, box_h,
            boxstyle="round,pad=0.12", facecolor='black', edgecolor='none',
            alpha=0.25, zorder=9
        )
        ax.add_patch(shadow)

        cat_box = FancyBboxPatch(
            (cx - box_w/2, cy - box_h/2), box_w, box_h,
            boxstyle="round,pad=0.12", facecolor=color, edgecolor='white',
            linewidth=1.8, alpha=0.9, zorder=10
        )
        ax.add_patch(cat_box)
        ax.text(cx, cy, cat_label, ha='center', va='center',
                fontsize=13, fontweight='bold', color='white', zorder=11,
                fontfamily=FONT_FAMILY)

        # Children
        child_start_y = cy - box_h/2 - 0.6
        child_spacing = 0.55

        for j, child in enumerate(children):
            child_y = child_start_y - j * child_spacing
            child_label = CHILD_ZH.get(child, child) if lang == "zh" else child

            # Vertical line from category down
            ax.plot([cx - 0.2, cx - 0.2], [cy - box_h/2, child_y],
                    color=color, lw=1.2, alpha=0.45, zorder=5)

            # Horizontal stub
            ax.plot([cx - 0.2, cx + 0.15], [child_y, child_y],
                    color=color, lw=1.2, alpha=0.45, zorder=5)

            # Bullet
            ax.plot(cx + 0.25, child_y, 's', color=color, markersize=4, zorder=6)

            # Child text with subtle background
            text_w = max(len(child_label) * 0.155, 1.5)
            text_bg = FancyBboxPatch(
                (cx + 0.45, child_y - 0.19), text_w, 0.38,
                boxstyle="round,pad=0.05", facecolor=color, edgecolor='none',
                alpha=0.08, zorder=5
            )
            ax.add_patch(text_bg)

            ax.text(cx + 0.55, child_y, child_label, ha='left', va='center',
                    fontsize=9.5, color='#E6EDF3', zorder=11,
                    fontfamily=FONT_FAMILY)

    # Bottom stats bar
    if lang == "en":
        stats = "327+ Exported Functions  ·  976+ Test Cases  ·  44 Platform Binaries  ·  11 Feature Categories  ·  66+ Sub-Features"
    else:
        stats = "327+ 导出函数  ·  976+ 测试用例  ·  44 平台二进制  ·  11 功能类别  ·  66+ 子功能"

    # Stats background
    stats_bg = FancyBboxPatch(
        (fig_width/2 - 8, 0.2), 16, 0.6,
        boxstyle="round,pad=0.1", facecolor='#21262D', edgecolor='#30363D',
        linewidth=1, alpha=0.8, zorder=8
    )
    ax.add_patch(stats_bg)
    ax.text(fig_width / 2, 0.5, stats, ha='center', va='center',
            fontsize=10, color='#8B949E', zorder=11,
            fontfamily=FONT_FAMILY)

    # Title watermark
    title_text = "cpe-skills Feature Mind Map" if lang == "en" else "cpe-skills 功能思维导图"
    ax.text(fig_width / 2, fig_height + 0.15, title_text,
            ha='center', va='bottom', fontsize=11, color='#484F58', style='italic',
            zorder=11, fontfamily=FONT_FAMILY)

    plt.tight_layout(pad=0.5)
    plt.savefig(output_path, dpi=150, bbox_inches='tight',
                facecolor='#0D1117', edgecolor='none')
    plt.close()
    print(f"Saved feature tree to {output_path}")


if __name__ == "__main__":
    draw_feature_tree("docs/images/feature_tree_en.png", lang="en")
    draw_feature_tree("docs/images/feature_tree_zh.png", lang="zh")

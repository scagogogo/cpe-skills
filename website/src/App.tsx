import { Layout, Typography, Button, Row, Col, Card, Tag, Space, Divider } from 'antd';
import {
  SafetyCertificateOutlined,
  ApiOutlined,
  CodeOutlined,
  ThunderboltOutlined,
  RocketOutlined,
  BranchesOutlined,
  DatabaseOutlined,
  SearchOutlined,
  ToolOutlined,
  CloudServerOutlined,
  GithubOutlined,
  CheckCircleOutlined,
  GlobalOutlined,
} from '@ant-design/icons';
import './App.css';

const { Header, Content, Footer } = Layout;
const { Title, Paragraph, Text } = Typography;

const features = [
  {
    icon: <SearchOutlined style={{ fontSize: 32, color: '#1677ff' }} />,
    title: 'CPE 解析与格式化',
    description: '自动识别 CPE 2.2 URI 和 2.3 Formatted String 格式，支持双向转换、WFN 绑定/解绑、特殊字符转义。',
    tags: ['CPE 2.2', 'CPE 2.3', 'WFN', 'FS', 'URI'],
  },
  {
    icon: <SafetyCertificateOutlined style={{ fontSize: 32, color: '#52c41a' }} />,
    title: 'NISTIR 7696 匹配',
    description: '完整实现 NIST 标准名称匹配语义：exact、subset、superset、disjoint，支持模糊匹配、正则匹配、版本范围。',
    tags: ['NISTIR 7696', '模糊匹配', '正则', '批量'],
  },
  {
    icon: <ToolOutlined style={{ fontSize: 32, color: '#faad14' }} />,
    title: 'CPE 生成与构建',
    description: '从产品信息、模板、模糊输入自动生成 CPE，提供 Fluent Builder API 和随机生成器。',
    tags: ['Builder', '模板', '模糊生成'],
  },
  {
    icon: <DatabaseOutlined style={{ fontSize: 32, color: '#722ed1' }} />,
    title: '漏洞关联',
    description: '多源漏洞数据集成：NVD、OSV、EPSS 概率评分、CISA KEV 已知被利用漏洞，一站式查询与关联。',
    tags: ['NVD', 'OSV', 'EPSS', 'KEV'],
  },
  {
    icon: <BranchesOutlined style={{ fontSize: 32, color: '#13c2c2' }} />,
    title: 'SBOM 与 PURL',
    description: '生成 CycloneDX / SPDX 格式 SBOM，CPE ↔ PURL 双向映射，依赖图分析，manifest 解析。',
    tags: ['CycloneDX', 'SPDX', 'PURL', '依赖图'],
  },
  {
    icon: <ThunderboltOutlined style={{ fontSize: 32, color: '#eb2f96' }} />,
    title: '风险评分与优先级',
    description: '综合 EPSS 概率、KEV 状态、可达性分析、VEX 排除，智能计算漏洞风险优先级。',
    tags: ['风险评分', '可达性', 'VEX', '优先级'],
  },
];

const integrationPaths = [
  {
    icon: <RocketOutlined style={{ fontSize: 40, color: '#1677ff' }} />,
    title: 'SKILLS',
    subtitle: 'AI/LLM 一键接入',
    description: '添加到 Claude Code skills 配置即可使用，AI 自动理解 CPE 语义并调用工具。',
    code: 'https://github.com/scagogogo/cpe-skills',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 40, color: '#52c41a' }} />,
    title: 'Go SDK',
    subtitle: '类型安全的 Go API',
    description: '完整的 Go 包，类型安全、零外部依赖（除 NVD API），支持所有核心功能。',
    code: 'go get github.com/scagogogo/cpe-skills',
  },
  {
    icon: <CloudServerOutlined style={{ fontSize: 40, color: '#faad14' }} />,
    title: 'CLI',
    subtitle: '命令行工具',
    description: '功能完整的 CLI，支持解析、匹配、搜索、字典操作，开箱即用。',
    code: 'go install github.com/scagogogo/cpe-skills/cmd/cpe@latest',
  },
  {
    icon: <ApiOutlined style={{ fontSize: 40, color: '#722ed1' }} />,
    title: 'MCP',
    subtitle: 'Model Context Protocol',
    description: '标准 MCP 服务器，任何支持 MCP 的 AI 工具都能直接调用 CPE 能力。',
    code: 'cpe mcp serve',
  },
];

const advantages = [
  '完整实现 NIST IR 7695/7696 CPE 标准',
  '91.6% 测试覆盖率，Race Detector 全量通过',
  '4 种接入方式：SKILLS / Go SDK / CLI / MCP',
  '多源漏洞数据集成（NVD + OSV + EPSS + KEV）',
  'CPE ↔ PURL 双向映射，SBOM 生成与分析',
  '智能风险评分与漏洞优先级排序',
  '零外部运行时依赖（纯 Go 实现）',
  '跨平台编译（Linux / macOS / Windows / MIPS）',
];

function App() {
  return (
    <Layout className="app-layout">
      {/* Header */}
      <Header className="app-header" style={{ position: 'fixed', top: 0, width: '100%', zIndex: 1000 }}>
        <div className="header-content">
          <div className="logo">
            <SafetyCertificateOutlined style={{ fontSize: 24, marginRight: 8 }} />
            <Text strong style={{ color: '#fff', fontSize: 18 }}>CPE Skills</Text>
          </div>
          <Space size="large">
            <a href="#features" className="nav-link">功能</a>
            <a href="#integration" className="nav-link">接入方式</a>
            <a href="#quickstart" className="nav-link">快速开始</a>
            <a href="https://github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer" className="nav-link">
              <GithubOutlined /> GitHub
            </a>
          </Space>
        </div>
      </Header>

      <Content className="app-content">
        {/* Hero Section */}
        <section className="hero-section">
          <div className="hero-content">
            <Title level={1} className="hero-title">
              CPE 全生命周期工具包
            </Title>
            <Paragraph className="hero-subtitle">
              从解析到漏洞管理，一站式解决 CPE (Common Platform Enumeration) 的所有问题。
              <br />
              完整实现 NIST 标准，支持 4 种接入方式，让网络安全工作不再繁琐。
            </Paragraph>
            <Space size="middle" style={{ marginTop: 32 }}>
              <Button type="primary" size="large" icon={<RocketOutlined />} href="#quickstart">
                快速开始
              </Button>
              <Button size="large" icon={<GithubOutlined />} href="https://github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer">
                GitHub
              </Button>
              <Button size="large" icon={<GlobalOutlined />} href="https://pkg.go.dev/github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer">
                Go Doc
              </Button>
            </Space>
            <div className="hero-tags" style={{ marginTop: 24 }}>
              <Space wrap>
                <Tag color="blue">CPE 2.2 / 2.3</Tag>
                <Tag color="green">NISTIR 7696</Tag>
                <Tag color="orange">NVD + OSV</Tag>
                <Tag color="purple">SBOM</Tag>
                <Tag color="cyan">EPSS + KEV</Tag>
                <Tag color="magenta">MCP</Tag>
              </Space>
            </div>
          </div>
        </section>

        {/* Problem Section */}
        <section className="problem-section">
          <div className="section-container">
            <Title level={2} className="section-title">
              为什么需要 CPE Skills？
            </Title>
            <Paragraph className="section-description">
              CPE 是 NIST 标准的 IT 系统命名方案，是 CVE 漏洞匹配、SBOM 组件追踪和供应链安全的基石。
              <br />
              <Text strong>但使用 CPE 很难：</Text>
            </Paragraph>
            <Row gutter={[24, 24]} style={{ marginTop: 32 }}>
              <Col xs={24} md={12}>
                <Card className="problem-card" hoverable>
                  <Title level={4}>😰 没有它时的痛点</Title>
                  <ul className="problem-list">
                    <li>CPE 有两种不兼容的格式（2.2 URI vs 2.3 格式化字符串）</li>
                    <li>WFN 绑定规则复杂，特殊字符转义容易出错</li>
                    <li>名称匹配需要理解 NISTIR 7696 关系语义</li>
                    <li>漏洞关联需要多源数据（NVD、OSV、EPSS、KEV）</li>
                    <li>SBOM 生成需要 CPE ↔ PURL 桥接</li>
                    <li>风险优先级需要整合 EPSS、KEV、可达性</li>
                  </ul>
                </Card>
              </Col>
              <Col xs={24} md={12}>
                <Card className="solution-card" hoverable>
                  <Title level={4}>✨ 有了它之后</Title>
                  <ul className="solution-list">
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> 一行代码自动识别并解析任意 CPE 格式</li>
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> WFN / FS / URI 自动双向转换</li>
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> NIST 标准匹配 + 模糊/正则/批量扩展</li>
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> 多源漏洞数据一站式查询与关联</li>
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> CPE ↔ PURL 映射 + SBOM 生成与分析</li>
                    <li><CheckCircleOutlined style={{ color: '#52c41a' }} /> 智能风险评分与漏洞优先级排序</li>
                  </ul>
                </Card>
              </Col>
            </Row>
          </div>
        </section>

        {/* Features Section */}
        <section id="features" className="features-section">
          <div className="section-container">
            <Title level={2} className="section-title">
              核心能力
            </Title>
            <Row gutter={[24, 24]}>
              {features.map((feature, index) => (
                <Col xs={24} md={12} lg={8} key={index}>
                  <Card className="feature-card" hoverable>
                    <div className="feature-icon">{feature.icon}</div>
                    <Title level={4}>{feature.title}</Title>
                    <Paragraph>{feature.description}</Paragraph>
                    <div className="feature-tags">
                      <Space wrap>
                        {feature.tags.map((tag) => (
                          <Tag key={tag}>{tag}</Tag>
                        ))}
                      </Space>
                    </div>
                  </Card>
                </Col>
              ))}
            </Row>
          </div>
        </section>

        {/* Integration Section */}
        <section id="integration" className="integration-section">
          <div className="section-container">
            <Title level={2} className="section-title">
              四种接入方式
            </Title>
            <Paragraph className="section-description">
              无论你是 AI 工具用户、Go 开发者、命令行爱好者还是 MCP 生态用户，都能零门槛接入。
            </Paragraph>
            <Row gutter={[24, 24]} style={{ marginTop: 32 }}>
              {integrationPaths.map((path, index) => (
                <Col xs={24} md={12} lg={6} key={index}>
                  <Card className="integration-card" hoverable>
                    <div className="integration-icon">{path.icon}</div>
                    <Title level={4}>{path.title}</Title>
                    <Text type="secondary" style={{ display: 'block', marginBottom: 8 }}>{path.subtitle}</Text>
                    <Paragraph>{path.description}</Paragraph>
                    <div className="code-block">
                      <code>{path.code}</code>
                    </div>
                  </Card>
                </Col>
              ))}
            </Row>
          </div>
        </section>

        {/* Quick Start Section */}
        <section id="quickstart" className="quickstart-section">
          <div className="section-container">
            <Title level={2} className="section-title">
              快速开始
            </Title>
            <Row gutter={[24, 24]}>
              <Col xs={24} md={8}>
                <Card className="quickstart-card" hoverable>
                  <Title level={4}>🤖 SKILLS 接入</Title>
                  <Paragraph>添加到 Claude Code skills 配置：</Paragraph>
                  <div className="code-block">
                    <code>https://github.com/scagogogo/cpe-skills</code>
                  </div>
                </Card>
              </Col>
              <Col xs={24} md={8}>
                <Card className="quickstart-card" hoverable>
                  <Title level={4}>📦 Go SDK</Title>
                  <Paragraph>安装并使用 Go 包：</Paragraph>
                  <div className="code-block">
                    <pre>{`go get github.com/scagogogo/cpe-skills

import cpeskills "github.com/scagogogo/cpe-skills"

c, _ := cpeskills.ParseCpe23(
  "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"
)`}</pre>
                  </div>
                </Card>
              </Col>
              <Col xs={24} md={8}>
                <Card className="quickstart-card" hoverable>
                  <Title level={4}>💻 CLI 使用</Title>
                  <Paragraph>安装命令行工具：</Paragraph>
                  <div className="code-block">
                    <pre>{`go install github.com/scagogogo/cpe-skills/cmd/cpe@latest

# 解析 CPE
cpe parse "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"

# 匹配 CPE
cpe match "cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*" \\
  "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"`}</pre>
                  </div>
                </Card>
              </Col>
            </Row>
          </div>
        </section>

        {/* Advantages Section */}
        <section className="advantages-section">
          <div className="section-container">
            <Title level={2} className="section-title">
              为什么选择 CPE Skills？
            </Title>
            <Row gutter={[24, 24]}>
              {advantages.map((adv, index) => (
                <Col xs={24} md={12} lg={6} key={index}>
                  <div className="advantage-item">
                    <CheckCircleOutlined style={{ fontSize: 24, color: '#52c41a', marginBottom: 8 }} />
                    <Text>{adv}</Text>
                  </div>
                </Col>
              ))}
            </Row>
          </div>
        </section>

        {/* CTA Section */}
        <section className="cta-section">
          <div className="section-container" style={{ textAlign: 'center' }}>
            <Title level={2}>现在就开始使用 CPE Skills</Title>
            <Paragraph style={{ fontSize: 16, color: 'rgba(255,255,255,0.8)' }}>
              让 CPE 不再是安全工作的障碍，而是你的利器。
            </Paragraph>
            <Space size="middle" style={{ marginTop: 24 }}>
              <Button type="primary" size="large" icon={<RocketOutlined />} href="https://github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer">
                访问 GitHub
              </Button>
              <Button size="large" ghost icon={<GlobalOutlined />} href="https://pkg.go.dev/github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer" style={{ color: '#fff', borderColor: '#fff' }}>
                查看 Go Doc
              </Button>
            </Space>
          </div>
        </section>
      </Content>

      {/* Footer */}
      <Footer className="app-footer">
        <div className="section-container">
          <Row gutter={[24, 24]}>
            <Col xs={24} md={8}>
              <Title level={5}>CPE Skills</Title>
              <Paragraph type="secondary">
                CPE 全生命周期工具包，从解析到漏洞管理，一站式解决方案。
              </Paragraph>
            </Col>
            <Col xs={24} md={8}>
              <Title level={5}>链接</Title>
              <Space direction="vertical">
                <a href="https://github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer">GitHub</a>
                <a href="https://pkg.go.dev/github.com/scagogogo/cpe-skills" target="_blank" rel="noreferrer">Go Doc</a>
                <a href="https://scagogogo.github.io/cpe-skills/" target="_blank" rel="noreferrer">文档</a>
              </Space>
            </Col>
            <Col xs={24} md={8}>
              <Title level={5}>标准</Title>
              <Space direction="vertical">
                <Text type="secondary">NIST IR 7695 — CPE Name Specification</Text>
                <Text type="secondary">NIST IR 7696 — CPE Name Matching</Text>
                <Text type="secondary">MITRE CVE — Vulnerability Enumeration</Text>
              </Space>
            </Col>
          </Row>
          <Divider style={{ borderColor: 'rgba(255,255,255,0.1)' }} />
          <Text type="secondary" style={{ textAlign: 'center', display: 'block' }}>
            © {new Date().getFullYear()} scagogogo — MIT License
          </Text>
        </div>
      </Footer>
    </Layout>
  );
}

export default App;

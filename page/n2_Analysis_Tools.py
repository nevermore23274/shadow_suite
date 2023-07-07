def code_analyzer():
    import streamlit as st
    from pycparser import parse_file, c_ast
    import pygraphviz as pgv
    import tempfile
    import os
    from io import BytesIO
    import re

    def remove_comments(text):
        return re.sub(re.compile("/\*.*?\*/", re.DOTALL), "", text) # remove all occurrences streamed comments (/*COMMENT */) from string

    # Inside the code_analyzer function, define the code_parser function
    def code_parser(c_code):
        import tempfile
        from io import BytesIO

        # Decode the bytes to string and filter out preprocessor lines
        c_code = '\n'.join(line for line in c_code.decode().split('\n') if not line.strip().startswith('#'))

        # Create a temporary file and write the content into it
        temp_file = tempfile.NamedTemporaryFile(delete=False, suffix='.c')
        with open(temp_file.name, 'w') as f:
            f.write(remove_comments(c_code)) # remove comments before writing to file

        ast = parse_file(temp_file.name, use_cpp=False)

        # Initialize an empty directed graph
        graph = pgv.AGraph(directed=True)

        # Set default node attributes
        graph.node_attr['color'] = 'white'  # Outline color of the node
        graph.node_attr['fillcolor'] = 'white'  # Fill color of the node
        graph.node_attr['fontcolor'] = 'black'  # Font color
        graph.node_attr['style'] = 'filled'  # Needed to apply the fill color

        # Set default edge attributes
        graph.edge_attr['color'] = 'white'  # Edge color

        # Define a visitor class and the visit_FuncDef method
        class FuncCallVisitor(c_ast.NodeVisitor):
            def visit_FuncCall(self, node):
                if node.name.name not in ("printf", "scanf"):  # Skip printf/scanf functions
                    graph.add_edge(node.name.name, parent_func)
                self.generic_visit(node)

        # Visit each function definition in the C program
        for ext in ast.ext:
            if isinstance(ext, c_ast.FuncDef):
                parent_func = ext.decl.name
                FuncCallVisitor().visit(ext)

        # Create a layout and render the graph
        graph.layout(prog='dot')
        output = BytesIO()
        graph.draw(output, prog='dot', format='svg')
        graph_svg = output.getvalue().decode('utf-8')

        # Modify SVG data
        graph_svg = graph_svg.replace('fill="white"', 'fill="black"')  # Set SVG background color to black
        graph_svg = graph_svg.replace('stroke="black"', 'stroke="white"')  # Set SVG text color to white

        # Close the temporary file and remove it
        temp_file.close()
        os.unlink(temp_file.name)

        return output.getvalue().decode('utf-8')

    # Still inside the code_analyzer function, handle the Streamlit app part
    uploaded_file = st.sidebar.file_uploader("Choose a C code file", type=['c', 'h'])

    if uploaded_file is not None:
        c_code = uploaded_file.read()
        graph_svg = code_parser(c_code)

        # Modify SVG data
        graph_svg = graph_svg.replace('fill="white"', 'fill="black"')  # Set SVG background color to black
        graph_svg = graph_svg.replace('stroke="black"', 'stroke="white"')  # Set SVG text color to white

        st.markdown(f'<div style="background-color: black; text-align: center">{graph_svg}</div>', unsafe_allow_html=True)
    else:
        st.error("Please upload a C code file.")

# Dictionary of subpage functions
page2_funcs = {
    "Code Analysis": code_analyzer
    #"Network Analysis": network_analysis,
    #"Subnet Calculator": subnet_calculator
}